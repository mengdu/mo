package mo

import (
	"context"
	"fmt"
	"os"
	"sync"
)

// Recorder is an interface for recording log messages.
type Recorder interface {
	// Log records a log message with the specified level and key-value pairs.
	Log(ctx context.Context, level Level, msg string, kv []KeyValue)
}

// New creates a new Logger instance with the specified context, recorder, and base key-value pairs.
func New(ctx context.Context, out Recorder, kv ...KeyValue) *Logger {
	return &Logger{
		ctx:   ctx,
		base:  kv,
		out:   out,
		level: LevelDebug,
		sprint: func(a ...interface{}) string {
			s := fmt.Sprintln(a...)
			return s[:len(s)-1] // remove \n at the end
		},
		sprintf: fmt.Sprintf,
		pool: &sync.Pool{
			New: func() interface{} {
				return new(KeyValue)
			},
		},
	}
}

// Logger is a logging client that provides methods for emitting log messages at different levels.
type Logger struct {
	ctx     context.Context                              // Context for the logger
	base    []KeyValue                                   // Base key-value pairs added to all log messages
	level   Level                                        // Minimum log level to emit
	out     Recorder                                     // Recorder for outputting log messages
	sprint  func(a ...interface{}) string                // Function for formatting arguments without a format string
	sprintf func(format string, a ...interface{}) string // Function for formatting arguments with a format string
	pool    *sync.Pool                                   // Pool for reusing KeyValue objects
}

// With creates a new Logger instance with the same configuration but using the specified context.
func (l Logger) With(ctx context.Context) *Logger {
	return &Logger{
		ctx:     ctx,
		base:    l.base,
		out:     l.out,
		level:   l.level,
		sprint:  l.sprint,
		sprintf: l.sprintf,
		pool:    l.pool,
	}
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
func (l *Logger) SetBase(kv ...KeyValue) {
	l.base = kv
}

// SetSprint sets the function used to format arguments without a format string.
func (l *Logger) SetSprint(sprint func(a ...interface{}) string) {
	l.sprint = sprint
}

// SetSprintf sets the function used to format arguments with a format string.
func (l *Logger) SetSprintf(sprintf func(format string, a ...interface{}) string) {
	l.sprintf = sprintf
}

// log is the internal method for logging messages at the specified level.
func (l Logger) log(level Level, msg string, kv ...KeyValue) {
	if !l.Enabled(level) {
		return
	}
	if l.out == nil {
		return
	}

	kvs := make([]KeyValue, 0, len(l.base)+len(kv))
	kvs = append(kvs, l.base...)
	kvs = append(kvs, kv...)

	for i, v := range kvs {
		if vv, ok := v.Value().(Valuer); ok {
			kvs[i][1] = vv(l.ctx)
		}
	}
	l.out.Log(l.ctx, level, msg, kvs)
}

// Debug logs a message at the debug level.
func (l Logger) Debug(a ...interface{}) {
	l.log(LevelDebug, l.sprint(a...))
}

// Info logs a message at the info level.
func (l Logger) Info(a ...interface{}) {
	l.log(LevelInfo, l.sprint(a...))
}

// Warn logs a message at the warn level.
func (l Logger) Warn(a ...interface{}) {
	l.log(LevelWarn, l.sprint(a...))
}

// Error logs a message at the error level.
func (l Logger) Error(a ...interface{}) {
	l.log(LevelError, l.sprint(a...))
}

// Fatal logs a message at the fatal level and exits the program.
func (l Logger) Fatal(a ...interface{}) {
	l.log(LevelFatal, l.sprint(a...))
	os.Exit(1)
}

// Debugf logs a formatted message at the debug level.
func (l Logger) Debugf(format string, a ...interface{}) {
	l.log(LevelDebug, l.sprintf(format, a...))
}

// Infof logs a formatted message at the info level.
func (l Logger) Infof(format string, a ...interface{}) {
	l.log(LevelInfo, l.sprintf(format, a...))
}

// Warnf logs a formatted message at the warn level.
func (l Logger) Warnf(format string, a ...interface{}) {
	l.log(LevelWarn, l.sprintf(format, a...))
}

// Errorf logs a formatted message at the error level.
func (l Logger) Errorf(format string, a ...interface{}) {
	l.log(LevelError, l.sprintf(format, a...))
}

// Fatalf logs a formatted message at the fatal level and exits the program.
func (l Logger) Fatalf(format string, a ...interface{}) {
	l.log(LevelFatal, l.sprintf(format, a...))
	os.Exit(1)
}

// Debugw logs a message with key-value pairs at the debug level.
func (l Logger) Debugw(msg string, kv ...KeyValue) {
	l.log(LevelDebug, msg, kv...)
}

// Infow logs a message with key-value pairs at the info level.
func (l Logger) Infow(msg string, kv ...KeyValue) {
	l.log(LevelInfo, msg, kv...)
}

// Warnw logs a message with key-value pairs at the warn level.
func (l Logger) Warnw(msg string, kv ...KeyValue) {
	l.log(LevelWarn, msg, kv...)
}

// Errorw logs a message with key-value pairs at the error level.
func (l Logger) Errorw(msg string, kv ...KeyValue) {
	l.log(LevelError, msg, kv...)
}

// Fatalw logs a message with key-value pairs at the fatal level and exits the program.
func (l Logger) Fatalw(msg string, kv ...KeyValue) {
	l.log(LevelFatal, msg, kv...)
	os.Exit(1)
}
