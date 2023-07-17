package mo

import (
	"errors"
	"log"
	"sync/atomic"
	"testing"
	"time"
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

type user struct {
	Name      string
	Email     string
	CreatedAt time.Time
}

var oneUser = &user{
	Name:      "Jane Doe",
	Email:     "jane@test.com",
	CreatedAt: time.Date(1980, 1, 1, 12, 0, 0, 0, time.UTC),
}
var meta = Meta{
	"int":     1,
	"ints":    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	"string":  "hello",
	"strings": []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
	"time":    time.Unix(0, 0),
	"times": []time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(2, 0),
		time.Unix(3, 0),
		time.Unix(4, 0),
		time.Unix(5, 0),
		time.Unix(6, 0),
		time.Unix(7, 0),
		time.Unix(8, 0),
		time.Unix(9, 0),
	},
	"user1": oneUser,
	"users": []*user{
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
	},
	"error": errors.New("fail"),
}

func BenchmarkNone(b *testing.B) {
	stream := &blackholeStream{}
	logger := New()
	logger.Level = LEVEL_NONE
	logger.Caller = false
	logger.DisableSprintfColor = true
	logger.Formater = &JsonForamter{}
	logger.Stdout = stream
	logger.Stderr = stream
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkDefault(b *testing.B) {
	stream := &blackholeStream{}
	logger := New()
	logger.Caller = false
	logger.ForceColor = true
	logger.Formater = &TextForamter{
		EnableTime: false,
	}
	logger.Stdout = stream
	logger.Stderr = stream
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkWith(b *testing.B) {
	stream := &blackholeStream{}
	logger := New()
	logger.Caller = false
	logger.Formater = &TextForamter{
		EnableTime: false,
	}
	logger.Stdout = stream
	logger.Stderr = stream
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.With(meta).Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkCaller(b *testing.B) {
	stream := &blackholeStream{}
	logger := New()
	logger.Caller = true
	logger.Formater = &TextForamter{
		EnableTime: false,
	}
	logger.Stdout = stream
	logger.Stderr = stream
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkFull(b *testing.B) {
	stream := &blackholeStream{}
	logger := New()
	logger.Caller = true
	logger.Meta = map[string]interface{}{
		"a": 1,
	}
	logger.ForceColor = true
	logger.Formater = &TextForamter{
		EnableTime:  true,
		EnableLevel: true,
	}
	logger.Stdout = stream
	logger.Stderr = stream
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.With(map[string]interface{}{
				"b": 1,
			}).Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkJsonForamter(b *testing.B) {
	stream := &blackholeStream{}
	logger := New()
	logger.Caller = false
	logger.Formater = &JsonForamter{}
	logger.Stdout = stream
	logger.Stderr = stream
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGoLog(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New(stream, "", log.Lshortfile)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Println("The quick brown fox jumps over the lazy dog")
		}
	})
	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
