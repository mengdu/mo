package mo

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

type blackholeStream struct {
	writeCount uint64
}

func (s *blackholeStream) WriteCount() uint64 {
	return atomic.LoadUint64(&s.writeCount)
}

func (s *blackholeStream) Write(p []byte) (int, error) {
	atomic.AddUint64(&s.writeCount, 1)
	return len(p), nil
}

type JSONLogger struct {
	encoder *json.Encoder
	mu      sync.Mutex
}

func (l *JSONLogger) Log(ctx context.Context, level Level, msg string, kv []KeyValue) {
	line := map[string]interface{}{
		"level": strings.ToLower(level.String()),
		"msg":   msg,
	}

	for _, v := range kv {
		line[v.Key()] = v.Value()
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	l.encoder.Encode(kv)
}

func BenchmarkDefault(b *testing.B) {
	stream := &blackholeStream{}
	std.SetRecorder(&stdRecorder{
		stdout: stream,
		stderr: stream,
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	})
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count %d != %d", stream.WriteCount(), b.N)
	}
}

func BenchmarkDefaultWithCaller(b *testing.B) {
	stream := &blackholeStream{}
	std.SetRecorder(&stdRecorder{
		stdout: stream,
		stderr: stream,
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	})
	SetBase(
		Value("ts", Timestamp("2006-01-02 15:04:05.000")),
		Value("caller", Caller(3)),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count %d != %d", stream.WriteCount(), b.N)
	}
}

func BenchmarkJson(b *testing.B) {
	stream := &blackholeStream{}
	out := &JSONLogger{
		encoder: json.NewEncoder(stream),
	}
	log := New(context.Background(),
		out,
		// Value("ts", Timestamp("2006-01-02 15:04:05.000")),
		// Value("caller", Caller(3)),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count %d != %d", stream.WriteCount(), b.N)
	}
}

func BenchmarkJsonWithCaller(b *testing.B) {
	stream := &blackholeStream{}
	out := &JSONLogger{
		encoder: json.NewEncoder(stream),
	}
	log := New(context.Background(),
		out,
		Value("ts", Timestamp("2006-01-02 15:04:05.000")),
		Value("caller", Caller(3)),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count %d != %d", stream.WriteCount(), b.N)
	}
}

func BenchmarkJsonWithCallerFull(b *testing.B) {
	stream := &blackholeStream{}
	out := &JSONLogger{
		encoder: json.NewEncoder(stream),
	}
	log := New(context.Background(),
		out,
		Value("ts", Timestamp("2006-01-02 15:04:05.000")),
		Value("caller", Caller(3)),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.With(context.Background()).Infow("The quick brown fox jumps over the lazy dog", Value("k1", "bar"), Value("k2", 123), Value("k3", 4.56), Value("k4", []int{1, 2, 3, 4, 5}))
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count %d != %d", stream.WriteCount(), b.N)
	}
}
