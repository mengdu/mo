package mo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
)

// Recorder is an interface for recording log messages.
type Recorder interface {
	// Log records a log message with the specified level and key-value pairs.
	Log(ctx context.Context, level Level, msg string, kv []Field)
}

type combine struct {
	Recorders []Recorder
}

func (c combine) Log(ctx context.Context, level Level, msg string, kv []Field) {
	for _, v := range c.Recorders {
		v.Log(ctx, level, msg, kv)
	}
}

func Combine(a ...Recorder) Recorder {
	return combine{Recorders: a}
}

// Ensure stdRecorder implements the Recorder interface.
var _ Recorder = (*stdRecorder)(nil)

// stdRecorder is a Recorder implementation that writes log messages to standard output and error.
type stdRecorder struct {
	stdout io.Writer  // Standard output writer
	stderr io.Writer  // Standard error writer
	mu     sync.Mutex // Mutex for concurrent access
	pool   *sync.Pool // Pool for reusing bytes.Buffer objects
}

// Log writes a log message to the appropriate output (stdout or stderr) based on the log level.
func (r *stdRecorder) Log(ctx context.Context, level Level, msg string, kv []Field) {
	buf := r.pool.Get().(*bytes.Buffer)
	defer r.pool.Put(buf)
	ts := ""
	caller := ""
	for _, v := range kv {
		if v.Key() == "ts" {
			ts = fmt.Sprint(v.Value())
		}
		if v.Key() == "caller" {
			if file, ok := v.Value().(string); ok {
				caller = file
			}
		}
	}

	if ts != "" {
		buf.WriteString("[")
		buf.WriteString(ts)
		buf.WriteString("]")
	}

	buf.WriteString("[")
	buf.WriteString(level.Abbr())
	buf.WriteString("] ")
	buf.WriteString(msg)

	i := 0
	for _, v := range kv {
		if v.Key() == "ts" || v.Key() == "caller" {
			continue
		}

		if i > 0 {
			buf.WriteString(", ")
		} else {
			buf.WriteString(" ")
		}
		fmt.Fprintf(buf, "%s=%v", v.Key(), v.Value())
		i++
	}
	if caller != "" {
		buf.WriteString(" ")
		buf.WriteString(caller)
	}
	buf.WriteByte('\n')
	defer buf.Reset()

	r.mu.Lock()
	defer r.mu.Unlock()
	if level >= LevelError {
		r.stderr.Write(buf.Bytes())
		return
	}
	r.stdout.Write(buf.Bytes())
}

// DefaultRecorder is the default Recorder implementation that writes to os.Stdout and os.Stderr.
var DefaultRecorder = &stdRecorder{
	stdout: os.Stdout,
	stderr: os.Stderr,
	pool: &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	},
}
