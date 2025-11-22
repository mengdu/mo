package record

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/mengdu/color"
	"github.com/mengdu/mo"
)

// Key constants for structured logging
const (
	KeyTimestamp = "ts"
	KeyCaller    = "caller"
	KeyMessage   = "msg"
	KeyLevel     = "level"
)

// Console is a simple console logger.
type Console struct {
	FilterEmptyField bool
	Stdout           io.Writer
	Stderr           io.Writer
	LevelType        string // "string", "abbr", "char"
	mu               sync.Mutex
	pool             *sync.Pool
}

// Log implements the Recorder interface.
func (c *Console) Log(ctx context.Context, level mo.Level, msg string, kv []mo.Field) {
	if c.pool == nil {
		c.pool = &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		}
	}

	buf := c.pool.Get().(*bytes.Buffer)
	defer c.pool.Put(buf)
	caller := ""
	ts := ""
	for _, v := range kv {
		if v.Key() == KeyTimestamp {
			ts = color.Dim().String(fmt.Sprintf("%v", v.Value()))
		}
		if v.Key() == KeyCaller {
			caller = v.Value().(string)
		}
	}

	if ts != "" {
		buf.WriteString("[")
		buf.WriteString(ts)
		buf.WriteString("]")
	}

	tag := ""
	if c.LevelType == "abbr" {
		tag = level.Abbr()
	} else if c.LevelType == "char" {
		tag = level.Char()
	} else {
		tag = level.String()
	}
	tag = fmt.Sprintf("[%s]", tag)
	switch level {
	case mo.LevelDebug:
		tag = color.BgGray().White().String(tag)
		msg = color.Gray().String(msg)
	case mo.LevelInfo:
		tag = color.BgBlue().White().String(tag)
	case mo.LevelWarn:
		tag = color.BgYellow().White().String(tag)
		msg = color.Yellow().String(msg)
	case mo.LevelError:
		tag = color.BgRed().White().String(tag)
		msg = color.Red().String(msg)
	case mo.LevelFatal:
		tag = color.BgRed().White().String(tag)
		msg = color.Red().String(msg)
	}

	buf.WriteString(tag)
	buf.WriteString(" ")
	buf.WriteString(msg)

	i := 0
	for _, v := range kv {
		if v.Key() == KeyTimestamp || v.Key() == KeyCaller {
			continue
		}

		val := fmt.Sprint(v.Value())
		// filter empty field
		if c.FilterEmptyField && val == "" {
			continue
		}

		if i > 0 {
			buf.WriteString(", ")
		} else {
			buf.WriteString(" ")
		}

		buf.WriteString(color.Gray().String(v.Key()))
		buf.WriteString("=")
		buf.WriteString(val)
		i++
	}
	if caller != "" {
		buf.WriteString(" ")
		buf.WriteString(color.Gray().Dim().String(caller))
	}
	buf.WriteByte('\n')
	defer buf.Reset()

	c.mu.Lock()
	defer c.mu.Unlock()
	if level >= mo.LevelError {
		if _, err := c.Stderr.Write(buf.Bytes()); err != nil {
			os.Stderr.WriteString(fmt.Sprintf("%v", err))
		}
		return
	}
	if _, err := c.Stdout.Write(buf.Bytes()); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v", err))
	}
}
