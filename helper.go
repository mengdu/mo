package mo

import (
	"context"
	"os"
)

type Helper struct {
	Logger *Logger
	ctx    context.Context
}

func New(ctx context.Context, logger *Logger) *Helper {
	return &Helper{ctx: ctx, Logger: logger}
}

func (h *Helper) With(ctx context.Context) *Helper {
	return &Helper{ctx: ctx, Logger: h.Logger}
}

// Debug logs a message at the debug level.
func (h Helper) Debug(a ...interface{}) {
	h.Logger.Log(h.ctx, LevelDebug, false, "", a, nil)
}

// Info logs a message at the info level.
func (h Helper) Info(a ...interface{}) {
	h.Logger.Log(h.ctx, LevelInfo, false, "", a, nil)
}

// Warn logs a message at the warn level.
func (h Helper) Warn(a ...interface{}) {
	h.Logger.Log(h.ctx, LevelWarn, false, "", a, nil)
}

// Error logs a message at the error level.
func (h Helper) Error(a ...interface{}) {
	h.Logger.Log(h.ctx, LevelError, false, "", a, nil)
}

// Fatal logs a message at the fatal level and exits the program.
func (h Helper) Fatal(a ...interface{}) {
	h.Logger.Log(h.ctx, LevelFatal, false, "", a, nil)
	os.Exit(1)
}

// Debugf logs a formatted message at the debug level.
func (h Helper) Debugf(format string, a ...interface{}) {
	h.Logger.Log(h.ctx, LevelDebug, true, format, a, nil)
}

// Infof logs a formatted message at the info level.
func (h Helper) Infof(format string, a ...interface{}) {
	h.Logger.Log(h.ctx, LevelInfo, true, format, a, nil)
}

// Warnf logs a formatted message at the warn level.
func (h Helper) Warnf(format string, a ...interface{}) {
	h.Logger.Log(h.ctx, LevelWarn, true, format, a, nil)
}

// Errorf logs a formatted message at the error level.
func (h Helper) Errorf(format string, a ...interface{}) {
	h.Logger.Log(h.ctx, LevelError, true, format, a, nil)
}

// Fatalf logs a formatted message at the fatal level and exits the program.
func (h Helper) Fatalf(format string, a ...interface{}) {
	h.Logger.Log(h.ctx, LevelFatal, true, format, a, nil)
	os.Exit(1)
}

// Debugw logs a message with key-value pairs at the debug level.
func (h Helper) Debugw(msg string, kv ...Field) {
	h.Logger.Log(h.ctx, LevelDebug, false, "", []interface{}{msg}, kv)
}

// Infow logs a message with key-value pairs at the info level.
func (h Helper) Infow(msg string, kv ...Field) {
	h.Logger.Log(h.ctx, LevelInfo, false, "", []interface{}{msg}, kv)
}

// Warnw logs a message with key-value pairs at the warn level.
func (h Helper) Warnw(msg string, kv ...Field) {
	h.Logger.Log(h.ctx, LevelWarn, false, "", []interface{}{msg}, kv)
}

// Errorw logs a message with key-value pairs at the error level.
func (h Helper) Errorw(msg string, kv ...Field) {
	h.Logger.Log(h.ctx, LevelError, false, "", []interface{}{msg}, kv)
}

// Fatalw logs a message with key-value pairs at the fatal level and exits the program.
func (h Helper) Fatalw(msg string, kv ...Field) {
	h.Logger.Log(h.ctx, LevelFatal, false, "", []interface{}{msg}, kv)
	os.Exit(1)
}
