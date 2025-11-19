package main

import (
	"context"

	"github.com/mengdu/mo"
)

type contextKey string

const requestIDKey contextKey = "request_id"

var TraceId = mo.Valuer(func(ctx context.Context) interface{} {
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		return id
	}
	return ""
})

func main() {
	mo.SetBase(
		mo.Value("ts", mo.Timestamp("15:04:05.000")),
		mo.Value("caller", mo.Caller(3)),
		mo.Value("trace.id", TraceId),
	)

	ctx := context.Background()
	ctx = context.WithValue(ctx, requestIDKey, "req-00a8c6cbfb48928d391aab5ef3676fe3417176c9")
	fields := []mo.Field{
		mo.Value("k1", 123),
		mo.Value("k2", true),
		mo.Value("k3", []interface{}{1, true, "test"}),
	}

	// Package-level functions - Debug
	mo.Debug("debug message", 123)
	mo.Debugf("debugf message %s", "test")
	mo.Debugw("debugw message", fields...)
	mo.Debugx(ctx, "debugx message", 123)
	mo.Debugfx(ctx, "debugfx message %s", "test")
	mo.Debugwx(ctx, "debugwx message", fields...)

	// Package-level functions - Info
	mo.Info("info message", 123)
	mo.Infof("infof message %s", "test")
	mo.Infow("infow message", fields...)
	mo.Infox(ctx, "infox message", 123)
	mo.Infofx(ctx, "infofx message %s", "test")
	mo.Infowx(ctx, "infowx message", fields...)

	// Package-level functions - Warn
	mo.Warn("warn message", 123)
	mo.Warnf("warnf message %s", "test")
	mo.Warnw("warnw message", fields...)
	mo.Warnx(ctx, "warnx message", 123)
	mo.Warnfx(ctx, "warnfx message %s", "test")
	mo.Warnwx(ctx, "warnwx message", fields...)

	// Package-level functions - Error
	mo.Error("error message", 123)
	mo.Errorf("errorf message %s", "test")
	mo.Errorw("errorw message", fields...)
	mo.Errorx(ctx, "errorx message", 123)
	mo.Errorfx(ctx, "errorfx message %s", "test")
	mo.Errorwx(ctx, "errorwx message", fields...)

	// Package-level functions - Fatal (commented out to avoid exit)
	// mo.Fatal("fatal message", 123)
	// mo.Fatalf("fatalf message %s", "test")
	// mo.Fatalw("fatalw message", fields...)
	// mo.Fatalx(ctx, "fatalx message", 123)
	// mo.Fatalfx(ctx, "fatalfx message %s", "test")
	// mo.Fatalwx(ctx, "fatalwx message", fields...)

	// Package-level with context
	ctx2 := context.WithValue(ctx, requestIDKey, "req-1234567890") // replace the request id
	mo.With(ctx2).Debug("debug message")
	mo.With(ctx2).Info("info message")
	mo.With(ctx2).Warn("warn message")
	mo.With(ctx2).Error("error message")
	// mo.With(ctx2).Fatal("fatal message")

	// Helper instance
	log := mo.New(context.Background(), mo.NewLogger(
		mo.DefaultRecorder,
		mo.Value("caller", mo.DefaultCaller),
		mo.Value("ts", mo.Timestamp("15:04:05.000")),
		mo.Value("trace.id", TraceId),
	))

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
