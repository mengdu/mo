package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/mengdu/color"
	"github.com/mengdu/fmtx"
	"github.com/mengdu/mo"
	"github.com/mengdu/mo/record"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/natefinch/lumberjack.v2"
)

type JSONLogger struct {
	encoder *json.Encoder
}

func (l *JSONLogger) Log(ctx context.Context, level mo.Level, msg string, kv []mo.Field) {
	line := map[string]interface{}{
		"level": strings.ToLower(level.String()),
		"msg":   msg,
	}

	for _, v := range kv {
		line[v.Key()] = v.Value()
	}

	l.encoder.Encode(line)
}

type ConsoleLogger struct {
	stdout io.Writer
	stderr io.Writer
	mu     sync.Mutex
	pool   *sync.Pool
}

func (c *ConsoleLogger) Log(ctx context.Context, level mo.Level, msg string, kv []mo.Field) {
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

	levelStr := fmt.Sprintf("[%s]", level.Char())
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

		if i > 0 {
			buf.WriteString(", ")
		} else {
			buf.WriteString(" ")
		}
		// fmt.Fprintf(buf, "%s=%v", color.Gray().String(v.Key()), v.Value())
		buf.WriteString(color.Gray().String(v.Key()))
		buf.WriteString("=")
		buf.WriteString(fmtx.String(v.Value()))
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
		c.stderr.Write(buf.Bytes())
		return
	}
	c.stdout.Write(buf.Bytes())
}

// TraceID returns a traceid valuer.
func TraceID() mo.Valuer {
	return func(ctx context.Context) interface{} {
		if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
			return span.TraceID().String()
		}
		return ""
	}
}

// SpanID returns a spanid valuer.
func SpanID() mo.Valuer {
	return func(ctx context.Context) interface{} {
		if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
			return span.SpanID().String()
		}
		return ""
	}
}

type mockSpan struct {
	trace.Span
	sc trace.SpanContext
}

func (s *mockSpan) SpanContext() trace.SpanContext {
	return s.sc
}

type LoggerOpts struct {
	Filename   string `mapstructure:"filename"`           // log file path
	Level      string `mapstructure:"level"`              // debug, info, warn, error, fatal
	Timestamp  string `mapstructure:"timestamp"`          // time.TimeOnly, time.RFC3339, time.Unix
	MaxSize    int    `mapstructure:"rolling.maxsize"`    // MB
	MaxBackups int    `mapstructure:"rolling.maxbackups"` // number of backups
	MaxAge     int    `mapstructure:"rolling.maxage"`     // days
	Compress   bool   `mapstructure:"rolling.compress"`   // disabled by default
}

func Init(ctx context.Context, opts *LoggerOpts) *mo.Helper {
	if err := os.MkdirAll(filepath.Dir(opts.Filename), 0755); err != nil {
		panic(err)
	}
	out := &lumberjack.Logger{
		Filename:   opts.Filename,
		MaxSize:    opts.MaxSize,    // 最大文件大小 1MB
		MaxBackups: opts.MaxBackups, // 备份数
		MaxAge:     opts.MaxAge,     // days
		Compress:   opts.Compress,   // 是否压缩
		LocalTime:  true,
	}

	consoleRecorder := &record.Console{
		FilterEmptyField: true,
		Stdout:           os.Stdout,
		Stderr:           os.Stderr,
		LevelType:        "char",
	}

	jsonRecorder := &record.JSON{
		Encoder: json.NewEncoder(out),
	}

	recorder := mo.Combine(consoleRecorder, jsonRecorder)
	// id, _ := os.Hostname()
	base := []mo.Field{
		mo.Value("ts", mo.Timestamp(opts.Timestamp)),
		mo.Value("caller", mo.Caller(3)),
		// mo.Value("service.id", id),
		mo.Value("trace.id", TraceID()),
		// mo.Value("span.id", SpanID()),
	}

	log := mo.NewLogger(recorder, base...)
	level := mo.ParseLevel(opts.Level)
	log.SetLevel(level)

	return mo.New(ctx, log)
}

func main() {
	ctx := context.Background()
	traceId, _ := trace.TraceIDFromHex("4bf92f3577b34da6a3ce929d0e0e4736")
	spanId, _ := trace.SpanIDFromHex("00f067aa0ba902b7")
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    traceId,
		SpanID:     spanId,
		TraceFlags: trace.FlagsSampled,
		Remote:     false,
	})

	ctx = trace.ContextWithSpan(ctx, &mockSpan{sc: sc})
	log := Init(ctx, &LoggerOpts{
		Filename:   "logs/app.log",
		Level:      "debug",
		Timestamp:  time.DateTime,
		MaxSize:    10,
		MaxBackups: 5,
		Compress:   true,
	})

	fields := []mo.Field{
		mo.Value("k1", 123),
		mo.Value("k2", true),
		mo.Value("k3", []interface{}{1, true, "test"}),
	}

	log.Debug("debug message")
	log.Debugf("debugf message %s", "test")
	log.Debugw("debugw message", fields...)
	log.Debugx(ctx, "debugx message")
	log.Debugfx(ctx, "debugfx message %s", "test")
	log.Debugwx(ctx, "debugwx message", fields...)

	log.Info("info message")
	log.Infof("infof message %s", "test")
	log.Infow("infow message", fields...)
	log.Infox(ctx, "infox message")
	log.Infofx(ctx, "infofx message %s", "test")
	log.Infowx(ctx, "infowx message", fields...)

	log.Warn("warn message")
	log.Warnf("warnf message %s", "test")
	log.Warnw("warnw message", fields...)
	log.Warnx(ctx, "warnx message")
	log.Warnfx(ctx, "warnfx message %s", "test")
	log.Warnwx(ctx, "warnwx message", fields...)

	log.Error("error message")
	log.Errorf("errorf message %s", "test")
	log.Errorw("errorw message", fields...)
	log.Errorx(ctx, "errorx message")
	log.Errorfx(ctx, "errorfx message %s", "test")
	log.Errorwx(ctx, "errorwx message", fields...)

	// log.Fatal("fatal message")
	// log.Fatalf("fatalf message %s", "test")
	// log.Fatalw("fatalw message", fields...)
	// log.Fatalx(ctx, "fatalx message")
	// log.Fatalfx(ctx, "fatalfx message %s", "test")
	// log.Fatalwx(ctx, "fatalwx message", fields...)
}
