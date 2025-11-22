package mo

import (
	"context"
	"os"
)

// std is the default Logger instance used by the package-level logging functions.
var std = New(context.Background(), NewLogger(DefaultRecorder))

// With returns a new Helper instance with the specified context.
func With(ctx context.Context) Helper {
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

// Debugf logs a formatted message at the debug level.
func Debugf(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelDebug, format, a...)
}

// Debugw logs a message with key-value pairs at the debug level.
func Debugw(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelDebug, msg, kv...)
}

// Debugx logs a message at the debug level with the given context.
func Debugx(ctx context.Context, a ...interface{}) {
	std.Logger.Print(ctx, LevelDebug, a...)
}

// Debugfx logs a formatted message at the debug level with the given context.
func Debugfx(ctx context.Context, format string, a ...interface{}) {
	std.Logger.Printf(ctx, LevelDebug, format, a...)
}

// Debugwx logs a message with key-value pairs at the debug level with the given context.
func Debugwx(ctx context.Context, msg string, kv ...Field) {
	std.Logger.Printw(ctx, LevelDebug, msg, kv...)
}

// Info logs a message at the info level.
func Info(a ...interface{}) {
	std.Logger.Print(std.ctx, LevelInfo, a...)
}

// Infof logs a formatted message at the info level.
func Infof(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelInfo, format, a...)
}

// Infow logs a message with key-value pairs at the info level.
func Infow(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelInfo, msg, kv...)
}

// Infox logs a message at the info level with the given context.
func Infox(ctx context.Context, a ...interface{}) {
	std.Logger.Print(ctx, LevelInfo, a...)
}

// Infofx logs a formatted message at the info level with the given context.
func Infofx(ctx context.Context, format string, a ...interface{}) {
	std.Logger.Printf(ctx, LevelInfo, format, a...)
}

// Infowx logs a message with key-value pairs at the info level with the given context.
func Infowx(ctx context.Context, msg string, kv ...Field) {
	std.Logger.Printw(ctx, LevelInfo, msg, kv...)
}

// Warn logs a message at the warn level.
func Warn(a ...interface{}) {
	std.Logger.Print(std.ctx, LevelWarn, a...)
}

// Warnf logs a formatted message at the warn level.
func Warnf(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelWarn, format, a...)
}

// Warnw logs a message with key-value pairs at the warn level.
func Warnw(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelWarn, msg, kv...)
}

// Warnx logs a message at the warn level with the given context.
func Warnx(ctx context.Context, a ...interface{}) {
	std.Logger.Print(ctx, LevelWarn, a...)
}

// Warnfx logs a formatted message at the warn level with the given context.
func Warnfx(ctx context.Context, format string, a ...interface{}) {
	std.Logger.Printf(ctx, LevelWarn, format, a...)
}

// Warnwx logs a message with key-value pairs at the warn level with the given context.
func Warnwx(ctx context.Context, msg string, kv ...Field) {
	std.Logger.Printw(ctx, LevelWarn, msg, kv...)
}

// Error logs a message at the error level.
func Error(a ...interface{}) {
	std.Logger.Print(std.ctx, LevelError, a...)
}

// Errorf logs a formatted message at the error level.
func Errorf(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelError, format, a...)
}

// Errorw logs a message with key-value pairs at the error level.
func Errorw(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelError, msg, kv...)
}

// Errorx logs a message at the error level with the given context.
func Errorx(ctx context.Context, a ...interface{}) {
	std.Logger.Print(ctx, LevelError, a...)
}

// Errorfx logs a formatted message at the error level with the given context.
func Errorfx(ctx context.Context, format string, a ...interface{}) {
	std.Logger.Printf(ctx, LevelError, format, a...)
}

// Errorwx logs a message with key-value pairs at the error level with the given context.
func Errorwx(ctx context.Context, msg string, kv ...Field) {
	std.Logger.Printw(ctx, LevelError, msg, kv...)
}

// Fatal logs a message at the fatal level and exits the program.
func Fatal(a ...interface{}) {
	std.Logger.Print(std.ctx, LevelFatal, a...)
	os.Exit(1)
}

// Fatalf logs a formatted message at the fatal level and exits the program.
func Fatalf(format string, a ...interface{}) {
	std.Logger.Printf(std.ctx, LevelFatal, format, a...)
	os.Exit(1)
}

// Fatalw logs a message with key-value pairs at the fatal level and exits the program.
func Fatalw(msg string, kv ...Field) {
	std.Logger.Printw(std.ctx, LevelFatal, msg, kv...)
	os.Exit(1)
}

// Fatalx logs a message at the fatal level with the given context and exits the program.
func Fatalx(ctx context.Context, a ...interface{}) {
	std.Logger.Print(ctx, LevelFatal, a...)
	os.Exit(1)
}

// Fatalfx logs a formatted message at the fatal level with the given context and exits the program.
func Fatalfx(ctx context.Context, format string, a ...interface{}) {
	std.Logger.Printf(ctx, LevelFatal, format, a...)
	os.Exit(1)
}

// Fatalwx logs a message with key-value pairs at the fatal level with the given context and exits the program.
func Fatalwx(ctx context.Context, msg string, kv ...Field) {
	std.Logger.Printw(ctx, LevelFatal, msg, kv...)
	os.Exit(1)
}
