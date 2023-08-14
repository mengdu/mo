package mo

import (
	"os"
)

var Std = &Logger{
	Stdout:   os.Stdout,
	Stderr:   os.Stderr,
	Level:    LEVEL_ALL,
	Formater: &TextForamter{},
	Caller:   false,
	// mu:       sync.Mutex{},
}

func Sprintf(format string, args ...any) string {
	return Std.Sprintf(format, args...)
}

func With(meta map[string]any) *Entry {
	return Std.With(meta)
}

func WithTag(tag string) *Entry {
	return Std.WithTag(tag)
}

func Error(s ...any) {
	Std.Error(s...)
}

func Errorf(format string, s ...any) {
	Std.Errorf(format, s...)
}

func Panic(s ...any) {
	Std.Panic(s...)
}

func Panicf(format string, s ...any) {
	Std.Panicf(format, s...)
}

func Warn(s ...any) {
	Std.Warn(s...)
}

func Warnf(format string, s ...any) {
	Std.Warnf(format, s...)
}

func Info(s ...any) {
	Std.Info(s...)
}

func Infof(format string, s ...any) {
	Std.Infof(format, s...)
}

func Log(s ...any) {
	Std.Log(s...)
}

func Logf(format string, s ...any) {
	Std.Logf(format, s...)
}

func Success(s ...any) {
	Std.Success(s...)
}

func Successf(format string, s ...any) {
	Std.Successf(format, s...)
}

func Debug(s ...any) {
	Std.Debug(s...)
}

func Debugf(format string, s ...any) {
	Std.Debugf(format, s...)
}
