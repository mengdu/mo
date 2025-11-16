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
func (h *Helper) With(ctx context.Context) *Helper {
	return &Helper{ctx: ctx, Logger: h.Logger}
}

// Debug logs a message at the debug level.
func (h Helper) Debug(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelDebug, a...)
}

// Info logs a message at the info level.
func (h Helper) Info(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelInfo, a...)
}

// Warn logs a message at the warn level.
func (h Helper) Warn(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelWarn, a...)
}

// Error logs a message at the error level.
func (h Helper) Error(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelError, a...)
}

// Fatal logs a message at the fatal level and exits the program.
func (h Helper) Fatal(a ...interface{}) {
	h.Logger.Print(h.ctx, LevelFatal, a...)
	os.Exit(1)
}

// Debugf logs a formatted message at the debug level.
func (h Helper) Debugf(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelDebug, format, a...)
}

// Infof logs a formatted message at the info level.
func (h Helper) Infof(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelInfo, format, a...)
}

// Warnf logs a formatted message at the warn level.
func (h Helper) Warnf(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelWarn, format, a...)
}

// Errorf logs a formatted message at the error level.
func (h Helper) Errorf(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelError, format, a...)
}

// Fatalf logs a formatted message at the fatal level and exits the program.
func (h Helper) Fatalf(format string, a ...interface{}) {
	h.Logger.Printf(h.ctx, LevelFatal, format, a...)
	os.Exit(1)
}

// Debugw logs a message with key-value pairs at the debug level.
func (h Helper) Debugw(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelDebug, msg, kv...)
}

// Infow logs a message with key-value pairs at the info level.
func (h Helper) Infow(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelInfo, msg, kv...)
}

// Warnw logs a message with key-value pairs at the warn level.
func (h Helper) Warnw(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelWarn, msg, kv...)
}

// Errorw logs a message with key-value pairs at the error level.
func (h Helper) Errorw(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelError, msg, kv...)
}

// Fatalw logs a message with key-value pairs at the fatal level and exits the program.
func (h Helper) Fatalw(msg string, kv ...Field) {
	h.Logger.Printw(h.ctx, LevelFatal, msg, kv...)
	os.Exit(1)
}
