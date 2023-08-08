package mo

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Entry struct {
	logger *Logger
	Meta   Meta
}

func (e *Entry) log(level Level, caller int, args ...any) {
	if !e.logger.IsEnableLevel(level) {
		return
	}
	filename := ""
	if e.logger.Caller {
		f := getCaller(caller)
		if f.File != "" {
			file := f.File
			if e.logger.RelativeFilePath && pwd != "" {
				p, _ := filepath.Rel(pwd, file)
				file = p
			}
			filename = fmt.Sprintf("%s:%d", file, f.Line)
		} else {
			filename = "(nil)"
		}
	}
	msg := fmt.Sprintln(args...)
	msg = msg[0 : len(msg)-1] // remove last \n
	log := &Record{
		Logger:   e.logger,
		At:       time.Now(),
		Tag:      e.logger.Tag,
		Level:    level,
		Message:  msg,
		Meta:     e.Meta,
		Filename: filename,
	}
	buf, err := e.logger.Formater.Format(log)
	if err != nil {
		fmt.Printf("Logger format error: %v\n", err)
	}

	e.logger.mu.Lock()
	defer e.logger.mu.Unlock()
	if log.Level == LEVEL_ERROR || log.Level == LEVEL_WARN {
		e.logger.Stderr.Write(buf)
	} else {
		e.logger.Stdout.Write(buf)
	}
}

func (e *Entry) With(meta map[string]any) *Entry {
	data := make(Meta, len(meta)+len(e.Meta))
	for k, v := range e.Meta {
		data[k] = v
	}
	for k, v := range meta {
		data[k] = v
	}

	return &Entry{
		logger: e.logger,
		Meta:   data,
	}
}

func (e *Entry) Error(s ...any) {
	e.log(LEVEL_ERROR, 4, s...)
}

func (e *Entry) Errorf(fotmat string, s ...any) {
	e.log(LEVEL_ERROR, 4, e.logger.Sprintf(fotmat, s...))
}

func (e *Entry) Panic(s ...any) {
	defer os.Exit(1)
	e.log(LEVEL_ERROR, 4, s...)
}

func (e *Entry) Panicf(fotmat string, s ...any) {
	defer os.Exit(1)
	e.log(LEVEL_ERROR, 4, e.logger.Sprintf(fotmat, s...))
}

func (e *Entry) Warn(s ...any) {
	e.log(LEVEL_WARN, 4, s...)
}

func (e *Entry) Warnf(fotmat string, s ...any) {
	e.log(LEVEL_WARN, 4, e.logger.Sprintf(fotmat, s...))
}

func (e *Entry) Info(s ...any) {
	e.log(LEVEL_INFO, 4, s...)
}

func (e *Entry) Infof(fotmat string, s ...any) {
	e.log(LEVEL_INFO, 4, e.logger.Sprintf(fotmat, s...))
}

func (e *Entry) Log(s ...any) {
	e.log(LEVEL_LOG, 4, s...)
}

func (e *Entry) Logf(fotmat string, s ...any) {
	e.log(LEVEL_LOG, 4, e.logger.Sprintf(fotmat, s...))
}

func (e *Entry) Success(s ...any) {
	e.log(LEVEL_SUCCESS, 4, s...)
}

func (e *Entry) Successf(fotmat string, s ...any) {
	e.log(LEVEL_SUCCESS, 4, e.logger.Sprintf(fotmat, s...))
}

func (e *Entry) Debug(s ...any) {
	e.log(LEVEL_DEBUG, 4, s...)
}

func (e *Entry) Debugf(fotmat string, s ...any) {
	e.log(LEVEL_DEBUG, 4, e.logger.Sprintf(fotmat, s...))
}
