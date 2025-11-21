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

type ConsoleRecorder struct {
	FilterEmptyField bool
	Stdout           io.Writer
	Stderr           io.Writer
	mu               sync.Mutex
	pool             *sync.Pool
}

func (c *ConsoleRecorder) Log(ctx context.Context, level mo.Level, msg string, kv []mo.Field) {
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
		if v.Key() == "ts" {
			ts = color.Dim().String(fmt.Sprintf("%v", v.Value()))
		}
		if v.Key() == "caller" {
			caller = v.Value().(string)
		}
	}

	if ts != "" {
		buf.WriteString("[")
		buf.WriteString(ts)
		buf.WriteString("]")
	}

	levelStr := fmt.Sprintf("[%s]", level.Abbr())
	switch level {
	case mo.LevelDebug:
		levelStr = color.BgGray().White().String(levelStr)
		msg = color.Gray().String(msg)
	case mo.LevelInfo:
		levelStr = color.BgBlue().White().String(levelStr)
	case mo.LevelWarn:
		levelStr = color.BgYellow().White().String(levelStr)
		msg = color.Yellow().String(msg)
	case mo.LevelError:
		levelStr = color.BgRed().White().String(levelStr)
		msg = color.Red().String(msg)
	case mo.LevelFatal:
		levelStr = color.BgRed().White().String(levelStr)
		msg = color.Red().String(msg)
	}

	buf.WriteString(levelStr)
	buf.WriteString(" ")
	buf.WriteString(msg)

	i := 0
	for _, v := range kv {
		if v.Key() == "ts" || v.Key() == "caller" {
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
