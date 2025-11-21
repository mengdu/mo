package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/mengdu/mo"
	"github.com/mengdu/mo/record"
	"github.com/natefinch/lumberjack"
	"go.opentelemetry.io/otel/trace"
)

type LoggerOpts struct {
	Filename   string `mapstructure:"filename"`
	Level      string `mapstructure:"level"`
	Timestamp  string `mapstructure:"timestamp"`
	MaxSize    int    `mapstructure:"rolling.maxsize"`    // MB
	MaxBackups int    `mapstructure:"rolling.maxbackups"` // 个数
	MaxAge     int    `mapstructure:"rolling.maxage"`     // days
	Compress   bool   `mapstructure:"rolling.compress"`
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

func Init(opts *LoggerOpts) *mo.Helper {
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

	consoleRecorder := &record.ConsoleRecorder{
		FilterEmptyField: true,
		Stdout:           os.Stdout,
		Stderr:           os.Stderr,
	}

	fileRecorder := &record.JSONRecorder{
		Encoder: json.NewEncoder(out),
	}

	recorder := mo.Combine(consoleRecorder, fileRecorder)
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

	return mo.New(context.Background(), log)
}

func main() {
	log := Init(&LoggerOpts{
		Filename:   "logs/app.log",
		Level:      "debug",
		Timestamp:  time.TimeOnly,
		MaxSize:    10,
		MaxBackups: 5,
		Compress:   true,
	})
	ctx := context.Background()
	fields := []mo.Field{
		mo.Value("k1", 123),
		mo.Value("k2", true),
		mo.Value("k3", []interface{}{1, true, "test"}),
	}

	// Helper methods - Debug
	log.Debug("debug message")
	log.Debugf("debugf message %s", "test")
	log.Debugw("debugw message", fields...)
	log.Debugx(ctx, "debugx message")
	log.Debugfx(ctx, "debugfx message %s", "test")
	log.Debugwx(ctx, "debugwx message", fields...)

	// Helper methods - Info
	log.Info("info message")
	log.Infof("infof message %s", "test")
	log.Infow("infow message", fields...)
	log.Infox(ctx, "infox message")
	log.Infofx(ctx, "infofx message %s", "test")
	log.Infowx(ctx, "infowx message", fields...)

	// Helper methods - Warn
	log.Warn("warn message")
	log.Warnf("warnf message %s", "test")
	log.Warnw("warnw message", fields...)
	log.Warnx(ctx, "warnx message")
	log.Warnfx(ctx, "warnfx message %s", "test")
	log.Warnwx(ctx, "warnwx message", fields...)

	// Helper methods - Error
	log.Error("error message")
	log.Errorf("errorf message %s", "test")
	log.Errorw("errorw message", fields...)
	log.Errorx(ctx, "errorx message")
	log.Errorfx(ctx, "errorfx message %s", "test")
	log.Errorwx(ctx, "errorwx message", fields...)

	// Helper methods - Fatal (commented out to avoid exit)
	// log.Fatal("fatal message")
	// log.Fatalf("fatalf message %s", "test")
	// log.Fatalw("fatalw message", fields...)
	// log.Fatalx(ctx, "fatalx message")
	// log.Fatalfx(ctx, "fatalfx message %s", "test")
	// log.Fatalwx(ctx, "fatalwx message", fields...)
}
