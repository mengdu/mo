package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/mengdu/mo"
	"gopkg.in/natefinch/lumberjack.v2"
)

type JSONLogger struct {
	encoder *json.Encoder
}

func (l *JSONLogger) Log(ctx context.Context, level mo.Level, msg string, kv []mo.KeyValue) {
	line := map[string]interface{}{
		"level": strings.ToLower(level.String()),
		"msg":   msg,
	}

	for _, v := range kv {
		line[v.Key()] = v.Value()
	}

	l.encoder.Encode(line)
}

func main() {
	filename := "./logs/app.log"
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		panic(err)
	}
	out := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1,     // 最大文件大小 1MB
		MaxBackups: 5,     // 备份数
		MaxAge:     28,    // days
		Compress:   false, // 是否压缩
		LocalTime:  false,
	}

	logger := &JSONLogger{
		encoder: json.NewEncoder(out),
	}

	mo.SetRecorder(logger)
	mo.SetBase(
		mo.Value("ts", mo.Timestamp("15:04:05.000")),
		mo.Value("caller", mo.Caller(4)),
		mo.Value("tag", "dev"),
	)

	for i := 0; i < 1000; i++ {
		demo()
	}
}

func demo() {
	mo.Debug("debug message")
	mo.Info("info message")
	mo.Warn("warn message")
	mo.Error("error message")
	// mo.Fatal("fatal message")

	mo.Debugf("debugf message %s", "test")
	mo.Infof("infof message %s", "test")
	mo.Warnf("warnf message %s", "test")
	mo.Errorf("errorf message %s", "test")
	// mo.Fatalf("fatalf message %s", "test")

	mo.Debugw("debugw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	mo.Infow("infow message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	mo.Warnw("warnw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	mo.Errorw("errorw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	// mo.Fatalw("fatalw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
}
