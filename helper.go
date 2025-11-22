package mo

import (
	"context"
	"os"
)

// Helper is a helper struct for logging.
type Helper struct {
	Logger *Logger
	ctx    context.Context
}

// New returns a new Helper with the given context and logger.
func New(ctx context.Context, logger *Logger) *Helper {
	return &Helper{ctx: ctx, Logger: logger}
}

// With returns a new Helper with the given context.
func (h Helper) With(ctx context.Context) *Helper {
	return &Helper{ctx: ctx, Logger: h.Logger}
}

// Debug logs a message at the debug level.
func (h Helper) Debug(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelDebug, a...)
}

// Debugf logs a formatted message at the debug level.
func (h Helper) Debugf(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelDebug, format, a...)
}

// Debugw logs a message with key-value pairs at the debug level.
func (h Helper) Debugw(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelDebug, msg, kv...)
}

// Debugx logs a message at the debug level with the given context.
func (h Helper) Debugx(ctx context.Context, a ...interface{}) {
	h.Logger.Print(ctx, LevelDebug, a...)
}

// Debugfx logs a formatted message at the debug level with the given context.
func (h Helper) Debugfx(ctx context.Context, format string, a ...interface{}) {
	h.Logger.Printf(ctx, LevelDebug, format, a...)
}

// Debugwx logs a message with key-value pairs at the debug level with the given context.
func (h Helper) Debugwx(ctx context.Context, msg string, kv ...Field) {
	h.Logger.Printw(ctx, LevelDebug, msg, kv...)
}

// Info logs a message at the info level.
func (h Helper) Info(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelInfo, a...)
}

// Infof logs a formatted message at the info level.
func (h Helper) Infof(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelInfo, format, a...)
}

// Infow logs a message with key-value pairs at the info level.
func (h Helper) Infow(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelInfo, msg, kv...)
}

// Infox logs a message at the info level with the given context.
func (h Helper) Infox(ctx context.Context, a ...interface{}) {
	h.Logger.Print(ctx, LevelInfo, a...)
}

// Infofx logs a formatted message at the info level with the given context.
func (h Helper) Infofx(ctx context.Context, format string, a ...interface{}) {
	h.Logger.Printf(ctx, LevelInfo, format, a...)
}

// Infowx logs a message with key-value pairs at the info level with the given context.
func (h Helper) Infowx(ctx context.Context, msg string, kv ...Field) {
	h.Logger.Printw(ctx, LevelInfo, msg, kv...)
}

// Warn logs a message at the warn level.
func (h Helper) Warn(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelWarn, a...)
}

// Warnf logs a formatted message at the warn level.
func (h Helper) Warnf(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelWarn, format, a...)
}

// Warnw logs a message with key-value pairs at the warn level.
func (h Helper) Warnw(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelWarn, msg, kv...)
}

// Warnx logs a message at the warn level with the given context.
func (h Helper) Warnx(ctx context.Context, a ...interface{}) {
	h.Logger.Print(ctx, LevelWarn, a...)
}

// Warnfx logs a formatted message at the warn level with the given context.
func (h Helper) Warnfx(ctx context.Context, format string, a ...interface{}) {
	h.Logger.Printf(ctx, LevelWarn, format, a...)
}

// Warnwx logs a message with key-value pairs at the warn level with the given context.
func (h Helper) Warnwx(ctx context.Context, msg string, kv ...Field) {
	h.Logger.Printw(ctx, LevelWarn, msg, kv...)
}

// Error logs a message at the error level.
func (h Helper) Error(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelError, a...)
}

// Errorf logs a formatted message at the error level.
func (h Helper) Errorf(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelError, format, a...)
}

// Errorw logs a message with key-value pairs at the error level.
func (h Helper) Errorw(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelError, msg, kv...)
}

// Errorx logs a message at the error level with the given context.
func (h Helper) Errorx(ctx context.Context, a ...interface{}) {
	h.Logger.Print(ctx, LevelError, a...)
}

// Errorfx logs a formatted message at the error level with the given context.
func (h Helper) Errorfx(ctx context.Context, format string, a ...interface{}) {
	h.Logger.Printf(ctx, LevelError, format, a...)
}

// Errorwx logs a message with key-value pairs at the error level with the given context.
func (h Helper) Errorwx(ctx context.Context, msg string, kv ...Field) {
	h.Logger.Printw(ctx, LevelError, msg, kv...)
}

// Fatal logs a message at the fatal level and exits the program.
func (h Helper) Fatal(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelFatal, a...)
	os.Exit(1)
}

// Fatalf logs a formatted message at the fatal level and exits the program.
func (h Helper) Fatalf(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelFatal, format, a...)
	os.Exit(1)
}

// Fatalw logs a message with key-value pairs at the fatal level and exits the program.
func (h Helper) Fatalw(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelFatal, msg, kv...)
	os.Exit(1)
}

// Fatalx logs a message at the fatal level with the given context and exits the program.
func (h Helper) Fatalx(ctx context.Context, a ...interface{}) {
	h.Logger.Print(ctx, LevelFatal, a...)
	os.Exit(1)
}

// Fatalfx logs a formatted message at the fatal level with the given context and exits the program.
func (h Helper) Fatalfx(ctx context.Context, format string, a ...interface{}) {
	h.Logger.Printf(ctx, LevelFatal, format, a...)
	os.Exit(1)
}

// Fatalwx logs a message with key-value pairs at the fatal level with the given context and exits the program.
func (h Helper) Fatalwx(ctx context.Context, msg string, kv ...Field) {
	h.Logger.Printw(ctx, LevelFatal, msg, kv...)
	os.Exit(1)
}
