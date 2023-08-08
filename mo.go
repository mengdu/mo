package mo

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/mattn/go-isatty"
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
	Tag      string    `json:"tag"`
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
	RelativeFilePath    bool
	ForceColor          bool
	DisableColor        bool
	DisableSprintfColor bool
	Tag                 string
	Level               Level
	mu                  sync.Mutex
	entryPool           sync.Pool
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

func (l *Logger) EnableColor() bool {
	enable := l.ForceColor || isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
	switch force, ok := os.LookupEnv("FORCE_COLOR"); {
	case ok && force != "0":
		enable = true
	}
	return enable && !l.DisableColor
}

func (l *Logger) Sprintf(format string, args ...any) string {
	if !l.EnableColor() || l.DisableSprintfColor {
		return fmt.Sprintf(format, args...)
	}
	b := newBuffer()
	defer releaseBuffer(b)
	end := len(format)
	for i := 0; i < end; {
		si := i
		for i < end && format[i] != '%' {
			i++
		}
		if i > si {
			b.WriteString(format[si:i])
		}
		if i >= end {
			break
		}
		pi := i
		i++ // skip %
	placeholder:
		for ; i < end; i++ {
			c := format[i]
			if c == '+' || c == '-' || c == '#' || c == ' ' || '0' <= c && c <= '9' {
				continue
			}
			pt := format[pi : i+1]
			switch c {
			case 'v':
				b.WriteString(color(pt, "93", "39"))
			case 'T', 'p':
				b.WriteString(color(pt, "34", "39"))
			case 't':
				b.WriteString(color(pt, "33", "39"))
			case 'd', 'e', 'E', 'o', 'x', 'X', 'b', 'f', 'F', 'g', 'G':
				b.WriteString(color(pt, "34", "39"))
			case 's', 'c', 'U', 'q':
				b.WriteString(color(pt, "32", "39"))
			default:
				b.WriteString(pt)
				i++
				break placeholder
			}
			i++
			break
		}
	}
	return fmt.Sprintf(b.String(), args...)
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
	l.log(LEVEL_ERROR, l.Sprintf(fotmat, s...))
}

func (l *Logger) Warn(s ...any) {
	l.log(LEVEL_WARN, s...)
}

func (l *Logger) Warnf(fotmat string, s ...any) {
	l.log(LEVEL_WARN, l.Sprintf(fotmat, s...))
}

func (l *Logger) Info(s ...any) {
	l.log(LEVEL_INFO, s...)
}

func (l *Logger) Infof(fotmat string, s ...any) {
	l.log(LEVEL_INFO, l.Sprintf(fotmat, s...))
}

func (l *Logger) Log(s ...any) {
	l.log(LEVEL_LOG, s...)
}

func (l *Logger) Logf(fotmat string, s ...any) {
	l.log(LEVEL_LOG, l.Sprintf(fotmat, s...))
}

func (l *Logger) Success(s ...any) {
	l.log(LEVEL_SUCCESS, s...)
}

func (l *Logger) Successf(fotmat string, s ...any) {
	l.log(LEVEL_SUCCESS, l.Sprintf(fotmat, s...))
}

func (l *Logger) Debug(s ...any) {
	l.log(LEVEL_DEBUG, s...)
}

func (l *Logger) Debugf(fotmat string, s ...any) {
	l.log(LEVEL_DEBUG, l.Sprintf(fotmat, s...))
}
