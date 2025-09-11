package mo

import (
	"context"
	"os"
)

// std is the default Logger instance used by the package-level logging functions.
var std = New(context.Background(), NewLogger(DefaultRecorder))

// With returns a new Helper instance with the specified context.
func With(ctx context.Context) *Helper {
	return std.With(ctx)
}

// Enabled returns whether logging at the specified level is enabled for the default logger.
func Enabled(level Level) bool {
	return std.Logger.Enabled(level)
}

// SetRecorder sets the recorder for the default logger.
func SetRecorder(out Recorder) {
	std.Logger.SetRecorder(out)
}

// SetLevel sets the log level for the default logger.
func SetLevel(level Level) {
	std.Logger.SetLevel(level)
}

// SetBase sets the base key-value pairs for the default logger.
func SetBase(kv ...Field) {
	std.Logger.SetBase(kv...)
}

// Debug logs a message at the debug level.
func Debug(a ...interface{}) {
	std.Logger.Log(std.ctx, LevelDebug, false, "", a, nil)
}

// Info logs a message at the info level.
func Info(a ...interface{}) {
	std.Logger.Log(std.ctx, LevelInfo, false, "", a, nil)
}

// Warn logs a message at the warn level.
func Warn(a ...interface{}) {
	std.Logger.Log(std.ctx, LevelWarn, false, "", a, nil)
}

// Error logs a message at the error level.
func Error(a ...interface{}) {
	std.Logger.Log(std.ctx, LevelError, false, "", a, nil)
}

// Fatal logs a message at the fatal level and exits the program.
func Fatal(a ...interface{}) {
	std.Logger.Log(std.ctx, LevelFatal, false, "", a, nil)
	os.Exit(1)
}

// Debugf logs a formatted message at the debug level.
func Debugf(format string, a ...interface{}) {
	std.Logger.Log(std.ctx, LevelDebug, true, format, a, nil)
}

// Infof logs a formatted message at the info level.
func Infof(format string, a ...interface{}) {
	std.Logger.Log(std.ctx, LevelInfo, true, format, a, nil)
}

// Warnf logs a formatted message at the warn level.
func Warnf(format string, a ...interface{}) {
	std.Logger.Log(std.ctx, LevelWarn, true, format, a, nil)
}

// Errorf logs a formatted message at the error level.
func Errorf(format string, a ...interface{}) {
	std.Logger.Log(std.ctx, LevelError, true, format, a, nil)
}

// Fatalf logs a formatted message at the fatal level and exits the program.
func Fatalf(format string, a ...interface{}) {
	std.Logger.Log(std.ctx, LevelFatal, true, format, a, nil)
	os.Exit(1)
}

// Debugw logs a message with key-value pairs at the debug level.
func Debugw(msg string, kv ...Field) {
	std.Logger.Log(std.ctx, LevelDebug, false, "", []interface{}{msg}, kv)
}

// Infow logs a message with key-value pairs at the info level.
func Infow(msg string, kv ...Field) {
	std.Logger.Log(std.ctx, LevelInfo, false, "", []interface{}{msg}, kv)
}

// Warnw logs a message with key-value pairs at the warn level.
func Warnw(msg string, kv ...Field) {
	std.Logger.Log(std.ctx, LevelWarn, false, "", []interface{}{msg}, kv)
}

// Errorw logs a message with key-value pairs at the error level.
func Errorw(msg string, kv ...Field) {
	std.Logger.Log(std.ctx, LevelError, false, "", []interface{}{msg}, kv)
}

// Fatalw logs a message with key-value pairs at the fatal level and exits the program.
func Fatalw(msg string, kv ...Field) {
	std.Logger.Log(std.ctx, LevelFatal, false, "", []interface{}{msg}, kv)
	os.Exit(1)
}
