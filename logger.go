package mo

import (
	"context"
	"fmt"
)

func sprint(a ...interface{}) string {
	s := fmt.Sprintln(a...)
	return s[:len(s)-1] // remove \n at the end
}

func sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// New creates a new Logger instance with the specified context, recorder, and base key-value pairs.
func NewLogger(out Recorder, kv ...Field) *Logger {
	return &Logger{
		base:  kv,
		level: LevelDebug,
		out:   out,
	}
}

// Logger is a logging client that provides methods for emitting log messages at different levels.
type Logger struct {
	base  []Field  // Base key-value pairs added to all log messages
	level Level    // Minimum log level to emit
	out   Recorder // Recorder for outputting log messages
}

// Enabled returns whether the specified log level is enabled.
func (l Logger) Enabled(level Level) bool {
	return level >= l.level
}

// SetLevel sets the minimum log level to emit.
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// SetRecorder sets the recorder used to output log messages.
func (l *Logger) SetRecorder(out Recorder) {
	l.out = out
}

// SetBase sets the base key-value pairs added to all log messages.
func (l *Logger) SetBase(kv ...Field) {
	l.base = kv
}

// log is the internal method for logging messages at the specified level.
// Deprecated: use Print, Printf or Printw instead.
func (l Logger) Log(ctx context.Context, level Level, formatting bool, format string, args []interface{}, kv []Field) {
	if !l.Enabled(level) {
		return
	}

	if l.out == nil {
		return
	}

	msg := ""
	if formatting {
		msg = sprintf(format, args...)
	} else {
		msg = sprint(args...)
	}

	kvs := make([]Field, 0, len(l.base)+len(kv))
	kvs = append(kvs, l.base...)
	kvs = append(kvs, kv...)

	for i, v := range kvs {
		if fn, ok := v.Value().(Valuer); ok {
			kvs[i][1] = fn(ctx)
		}
	}
	l.out.Log(ctx, level, msg, kvs)
}

// Print logs a message at the specified level.
func (l Logger) Print(ctx context.Context, level Level, a ...interface{}) {
	if !l.Enabled(level) || l.out == nil {
		return
	}

	kvs := append(make([]Field, 0, len(l.base)), l.base...)
	for i, v := range kvs {
		if fn, ok := v.Value().(Valuer); ok {
			kvs[i][1] = fn(ctx)
		}
	}

	l.out.Log(ctx, level, sprint(a...), kvs)
}

// Printf logs a formatted message at the specified level.
func (l Logger) Printf(ctx context.Context, level Level, format string, a ...interface{}) {
	if !l.Enabled(level) || l.out == nil {
		return
	}

	kvs := append(make([]Field, 0, len(l.base)), l.base...)
	for i, v := range kvs {
		if fn, ok := v.Value().(Valuer); ok {
			kvs[i][1] = fn(ctx)
		}
	}

	l.out.Log(ctx, level, sprintf(format, a...), kvs)
}

// Printw logs a message with key-value pairs at the specified level.
func (l Logger) Printw(ctx context.Context, level Level, msg string, kv ...Field) {
	if !l.Enabled(level) || l.out == nil {
		return
	}

	kvs := append(make([]Field, 0, len(l.base)+len(kv)), l.base...)
	kvs = append(kvs, kv...)
	for i, v := range kvs {
		if fn, ok := v.Value().(Valuer); ok {
			kvs[i][1] = fn(ctx)
		}
	}

	l.out.Log(ctx, level, msg, kvs)
}
