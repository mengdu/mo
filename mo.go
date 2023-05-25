package mo

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type any = interface{}
type Meta = map[string]any

var pwd, _ = os.Getwd()

type Formater interface {
	Format(log *Record) ([]byte, error)
}

type Record struct {
	Logger   *Logger   `json:"-"`
	At       time.Time `json:"at"`
	Level    Level     `json:"level"`
	Message  string    `json:"msg"`
	Meta     Meta      `json:"meta"`
	Filename string    `json:"file"`
}

type Logger struct {
	Stdout   io.Writer
	Stderr   io.Writer
	Formater Formater
	// log meta fields
	Meta Meta
	// Caller tracing
	Caller bool
	// Caller tracing file is relative path
	RelativeFilePath bool
	Level            Level
	mu               sync.Mutex
	entryPool        sync.Pool
}

func New() *Logger {
	return &Logger{
		Stdout:   os.Stdout,
		Stderr:   os.Stderr,
		Formater: &TextForamter{},
		Level:    LEVEL_ALL,
		mu:       sync.Mutex{},
	}
}

func (l *Logger) newEntry() *Entry {
	meta := make(Meta, len(l.Meta))
	for k, v := range l.Meta {
		meta[k] = v
	}
	if entry, ok := l.entryPool.Get().(*Entry); ok {
		entry.Meta = meta
		return entry
	}
	return &Entry{
		logger: l,
		Meta:   meta,
	}
}

func (l *Logger) releaseEntry(entry *Entry) {
	entry.Meta = make(Meta, 0)
	l.entryPool.Put(entry)
}

func (l *Logger) log(level Level, s ...any) {
	if !isEnableLevel(l.Level, level) {
		return
	}
	entry := l.newEntry()
	entry.log(level, 5, s...)
	l.releaseEntry(entry)
}

func (l *Logger) IsEnableLevel(level Level) bool {
	return isEnableLevel(l.Level, level)
}

func (l *Logger) With(meta map[string]any) *Entry {
	// entry := l.newEntry()
	// defer l.releaseEntry(entry)
	// return entry.With(meta)
	data := make(Meta, len(meta)+len(l.Meta))
	for k, v := range l.Meta {
		data[k] = v
	}
	for k, v := range meta {
		data[k] = v
	}

	return &Entry{
		logger: l,
		Meta:   data,
	}
}

func (l *Logger) Error(s ...any) {
	l.log(LEVEL_ERROR, s...)
}

func (l *Logger) Errorf(fotmat string, s ...any) {
	l.log(LEVEL_ERROR, fmt.Sprintf(fotmat, s...))
}

func (l *Logger) Warn(s ...any) {
	l.log(LEVEL_WARN, s...)
}

func (l *Logger) Warnf(fotmat string, s ...any) {
	l.log(LEVEL_WARN, fmt.Sprintf(fotmat, s...))
}

func (l *Logger) Info(s ...any) {
	l.log(LEVEL_INFO, s...)
}

func (l *Logger) Infof(fotmat string, s ...any) {
	l.log(LEVEL_INFO, fmt.Sprintf(fotmat, s...))
}

func (l *Logger) Log(s ...any) {
	l.log(LEVEL_LOG, s...)
}

func (l *Logger) Logf(fotmat string, s ...any) {
	l.log(LEVEL_LOG, fmt.Sprintf(fotmat, s...))
}

func (l *Logger) Success(s ...any) {
	l.log(LEVEL_SUCCESS, s...)
}

func (l *Logger) Successf(fotmat string, s ...any) {
	l.log(LEVEL_SUCCESS, fmt.Sprintf(fotmat, s...))
}

func (l *Logger) Debug(s ...any) {
	l.log(LEVEL_DEBUG, s...)
}

func (l *Logger) Debugf(fotmat string, s ...any) {
	l.log(LEVEL_DEBUG, fmt.Sprintf(fotmat, s...))
}
