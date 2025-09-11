package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/mengdu/color"
	"github.com/mengdu/fmtx"
	"github.com/mengdu/mo"
	"go.opentelemetry.io/otel/trace"
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
	// logger := &JSONLogger{
	// 	encoder: json.NewEncoder(os.Stdout),
	// }
	logger := &ConsoleLogger{
		stdout: os.Stdout,
		stderr: os.Stderr,
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
	id, _ := os.Hostname()
	log := mo.NewLogger(logger,
		mo.Value("ts", mo.Timestamp("2006-01-02 15:04:05.000")),
		mo.Value("caller", mo.Caller()),
		mo.Value("service.id", id),
		mo.Value("service.version", "v1.2.3"),
		mo.Value("trace.id", TraceID()),
		mo.Value("span.id", SpanID()),
	)
	l := mo.New(ctx, log)
	// l.Logger.SetLevel(mo.LevelError)
	// l.Logger.SetLevel(mo.LevelInfo)

	l.Debug("debug message")
	l.Info("info message")
	l.Warn("warn message")
	l.Error("error message")
	// l.Fatal("fatal message")

	l.Debugf("debugf message %s", "test")
	l.Infof("infof message %s", "test")
	l.Warnf("warnf message %s", "test")
	l.Errorf("errorf message %s", "test")
	// l.Fatalf("fatalf message %s", "test")

	l.Debugw("debugw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	l.Infow("infow message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	l.Warnw("warnw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	l.Errorw("errorw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	// l.Fatalw("fatalw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))

	l.Infow("replace ts, caller", mo.Value("ts", "xxx"), mo.Value("caller", "path-to-xxx.go:123"))
	l.With(context.Background()).Infow("test with context", mo.Value("k1", 123))
}
