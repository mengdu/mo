package mo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// discardWriter is a writer that discards all writes (similar to io.Discard but with counting)
type discardWriter struct {
	writeCount uint64
}

func (w *discardWriter) Write(p []byte) (int, error) {
	atomic.AddUint64(&w.writeCount, 1)
	return len(p), nil
}

func (w *discardWriter) WriteCount() uint64 {
	return atomic.LoadUint64(&w.writeCount)
}

// jsonRecorder is a Recorder implementation that outputs JSON
type jsonRecorder struct {
	encoder *json.Encoder
	mu      sync.Mutex
}

func newJSONRecorder(w io.Writer) *jsonRecorder {
	return &jsonRecorder{
		encoder: json.NewEncoder(w),
	}
}

func (r *jsonRecorder) Log(ctx context.Context, level Level, msg string, kv []Field) {
	line := map[string]interface{}{
		"level": strings.ToLower(level.String()),
		"msg":   msg,
	}

	for _, v := range kv {
		line[v.Key()] = v.Value()
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if err := r.encoder.Encode(line); err != nil {
		panic(err)
	}
}

func setupBenchmarkHelper() (*Helper, *discardWriter) {
	discard := &discardWriter{}
	recorder := &stdRecorder{
		stdout: discard,
		stderr: discard,
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
	logger := NewLogger(recorder)
	return New(context.Background(), logger), discard
}

type key string

const mykey = key("req-id")

var fields = []Field{
	Value("k1", 123),
	Value("k2", true),
	Value("k3", []interface{}{1, true, "false", map[string]interface{}{"a": 1}}),
	Value("k4", Valuer(func(ctx context.Context) interface{} {
		v, _ := ctx.Value(mykey).(string)
		return v
	})),
}

// Benchmark tests for Helper methods

func Benchmark_Info(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("test message")
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_Infof(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infof("test message %s", "test")
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_Infow(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infow("test message", fields...)
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_Infox(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	ctx := context.Background()
	ctx = context.WithValue(ctx, mykey, "123")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infox(ctx, "test message")
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_Infofx(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	ctx := context.Background()
	ctx = context.WithValue(ctx, mykey, "123")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infofx(ctx, "test message %s", "test")
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_Infowx(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	ctx := context.Background()
	ctx = context.WithValue(ctx, mykey, "123")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infowx(ctx, "test message", fields...)
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_WithCaller_Info(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	log.Logger.SetBase(Value("caller", DefaultCaller), Value("ts", Timestamp(time.RFC3339)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("test message")
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_WithCaller_Infof(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	log.Logger.SetBase(Value("caller", DefaultCaller), Value("ts", Timestamp(time.RFC3339)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infof("test message %s", "test")
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_WithCaller_Infow(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	log.Logger.SetBase(Value("caller", DefaultCaller), Value("ts", Timestamp(time.RFC3339)))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infow("test message", fields...)
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_WithCaller_Infox(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	log.Logger.SetBase(Value("caller", DefaultCaller), Value("ts", Timestamp(time.RFC3339)))
	ctx := context.Background()
	ctx = context.WithValue(ctx, mykey, "123")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infox(ctx, "test message")
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_WithCaller_Infofx(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	log.Logger.SetBase(Value("caller", DefaultCaller), Value("ts", Timestamp(time.RFC3339)))
	ctx := context.Background()
	ctx = context.WithValue(ctx, mykey, "123")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infofx(ctx, "test message %s", "test")
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_WithCaller_Infowx(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	log.Logger.SetBase(Value("caller", DefaultCaller), Value("ts", Timestamp(time.RFC3339)))
	ctx := context.Background()
	ctx = context.WithValue(ctx, mykey, "123")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Infowx(ctx, "test message", fields...)
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_With(b *testing.B) {
	log, discard := setupBenchmarkHelper()
	log.Logger.SetBase(Value("caller", DefaultCaller), Value("ts", Timestamp(time.RFC3339)))
	ctx := context.Background()
	ctx = context.WithValue(ctx, mykey, "123")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.With(ctx).Info("test message")
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}

func Benchmark_JSON(b *testing.B) {
	log, _ := setupBenchmarkHelper()
	log.Logger.SetBase(Value("caller", DefaultCaller), Value("ts", Timestamp(time.RFC3339)))
	discard := &discardWriter{}
	record := newJSONRecorder(discard)
	log.Logger.SetRecorder(record)
	ctx := context.Background()
	ctx = context.WithValue(ctx, mykey, "123")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// log.Info("test message")
			// log.Infof("test message %s", "test")
			// log.Infow("test message", fields...)
			// log.Infox(ctx, "test message")
			// log.Infofx(ctx, "test message %s", "test")
			log.Infowx(ctx, "test message", fields...)
		}
	})
	if discard.WriteCount() != uint64(b.N) {
		b.Errorf("expected %d writes, got %d", b.N, discard.WriteCount())
	}
}
