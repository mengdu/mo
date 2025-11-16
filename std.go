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
	std.Logger.Print(std.ctx, LevelDebug, a...)
}

// Info logs a message at the info level.
func Info(a ...interface{}) {
	std.Logger.Print(std.ctx, LevelInfo, a...)
}

// Warn logs a message at the warn level.
func Warn(a ...interface{}) {
	std.Logger.Print(std.ctx, LevelWarn, a...)
}

// Error logs a message at the error level.
func Error(a ...interface{}) {
	std.Logger.Print(std.ctx, LevelError, a...)
}

// Fatal logs a message at the fatal level and exits the program.
func Fatal(a ...interface{}) {
	std.Logger.Print(std.ctx, LevelFatal, a...)
	os.Exit(1)
}

// Debugf logs a formatted message at the debug level.
func Debugf(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelDebug, format, a...)
}

// Infof logs a formatted message at the info level.
func Infof(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelInfo, format, a...)
}

// Warnf logs a formatted message at the warn level.
func Warnf(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelWarn, format, a...)
}

// Errorf logs a formatted message at the error level.
func Errorf(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelError, format, a...)
}

// Fatalf logs a formatted message at the fatal level and exits the program.
func Fatalf(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelFatal, format, a...)
	os.Exit(1)
}

// Debugw logs a message with key-value pairs at the debug level.
func Debugw(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelDebug, msg, kv...)
}

// Infow logs a message with key-value pairs at the info level.
func Infow(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelInfo, msg, kv...)
}

// Warnw logs a message with key-value pairs at the warn level.
func Warnw(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelWarn, msg, kv...)
}

// Errorw logs a message with key-value pairs at the error level.
func Errorw(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelError, msg, kv...)
}

// Fatalw logs a message with key-value pairs at the fatal level and exits the program.
func Fatalw(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelFatal, msg, kv...)
	os.Exit(1)
}
