package main

import (
	"context"

	"github.com/mengdu/mo"
)

func main() {
	// mo.SetBase(
	// 	mo.Value("ts", mo.Timestamp("15:04:05.000")),
	// 	mo.Value("caller", mo.Caller(4)),
	// 	mo.Value("tag", "dev"),
	// )

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

	mo.With(context.Background()).Info("test with context")
}
