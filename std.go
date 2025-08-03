package mo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
)

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
func (r *stdRecorder) Log(ctx context.Context, level Level, msg string, kv []KeyValue) {
	buf := r.pool.Get().(*bytes.Buffer)
	defer r.pool.Put(buf)
	ts := ""
	caller := ""
	for _, v := range kv {
		if v.Key() == "ts" {
			ts = fmt.Sprint(v.Value())
		}
		if v.Key() == "caller" {
			caller = v.Value().(string)
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

// std is the default Logger instance used by the package-level logging functions.
var std = New(context.Background(), DefaultRecorder)

// Enabled returns whether the debug log level is enabled for the default logger.
func Enabled() bool {
	return std.Enabled(LevelDebug)
}

// SetRecorder sets the recorder for the default logger.
func SetRecorder(out Recorder) {
	std.SetRecorder(out)
}

// SetLevel sets the log level for the default logger.
func SetLevel(level Level) {
	std.SetLevel(level)
}

// SetBase sets the base key-value pairs for the default logger.
func SetBase(kv ...KeyValue) {
	std.SetBase(kv...)
}

// Debug logs a message at the debug level using the default logger.
func Debug(a ...interface{}) {
	std.Debug(a...)
}

// Info logs a message at the info level using the default logger.
func Info(a ...interface{}) {
	std.Info(a...)
}

// Warn logs a message at the warn level using the default logger.
func Warn(a ...interface{}) {
	std.Warn(a...)
}

// Error logs a message at the error level using the default logger.
func Error(a ...interface{}) {
	std.Error(a...)
}

// Fatal logs a message at the fatal level using the default logger and exits the program.
func Fatal(a ...interface{}) {
	std.Fatal(a...)
}

// Debugf logs a formatted message at the debug level using the default logger.
func Debugf(format string, a ...interface{}) {
	std.Debugf(format, a...)
}

// Infof logs a formatted message at the info level using the default logger.
func Infof(format string, a ...interface{}) {
	std.Infof(format, a...)
}

// Warnf logs a formatted message at the warn level using the default logger.
func Warnf(format string, a ...interface{}) {
	std.Warnf(format, a...)
}

// Errorf logs a formatted message at the error level using the default logger.
func Errorf(format string, a ...interface{}) {
	std.Errorf(format, a...)
}

// Fatalf logs a formatted message at the fatal level using the default logger and exits the program.
func Fatalf(format string, a ...interface{}) {
	std.Fatalf(format, a...)
}

// Debugw logs a message with key-value pairs at the debug level using the default logger.
func Debugw(msg string, kv ...KeyValue) {
	std.Debugw(msg, kv...)
}

// Infow logs a message with key-value pairs at the info level using the default logger.
func Infow(msg string, kv ...KeyValue) {
	std.Infow(msg, kv...)
}

// Warnw logs a message with key-value pairs at the warn level using the default logger.
func Warnw(msg string, kv ...KeyValue) {
	std.Warnw(msg, kv...)
}

// Errorw logs a message with key-value pairs at the error level using the default logger.
func Errorw(msg string, kv ...KeyValue) {
	std.Errorw(msg, kv...)
}

// Fatalw logs a message with key-value pairs at the fatal level using the default logger and exits the program.
func Fatalw(msg string, kv ...KeyValue) {
	std.Fatalw(msg, kv...)
}

func With(ctx context.Context) *Logger {
	log := std.With(ctx)
	base := make([]KeyValue, len(log.base))
	for i, v := range log.base {
		// replace caller
		if v.Key() == "caller" {
			base[i] = Value(v.Key(), Caller(3))
			continue
		}
		base[i] = v
	}
	log.SetBase(base...)
	return log
}
