package main

import (
	"context"

	"github.com/mengdu/mo"
)

func main() {
	mo.SetBase(
		mo.Value("ts", mo.Timestamp("15:04:05.000")),
		mo.Value("caller", mo.Caller()),
		mo.Value("tag", "dev"),
	)

	mo.Debug("debug message", 123)
	mo.Info("info message", 123)
	mo.Warn("warn message", 123)
	mo.Error("error message", 123)
	// mo.Fatal("fatal message", 123)

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

	mo.Infow("replace ts, caller", mo.Value("ts", "xxx"), mo.Value("caller", "path-to-xxx.go:123"))
	mo.With(context.Background()).Infow("test with context", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))

	log := mo.New(context.Background(), mo.NewLogger(
		mo.DefaultRecorder,
		mo.Value("caller", mo.Caller()),
		mo.Value("ts", mo.Timestamp("15:04:05.000")),
	))

	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")
	// log.Fatal("fatal message")
	log.Debugf("debugf message %s", "test")
	log.Infof("infof message %s", "test")
	log.Warnf("warnf message %s", "test")
	log.Errorf("errorf message %s", "test")
	// log.Fatalf("fatalf message %s", "test")
	log.Debugw("debugw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	log.Infow("infow message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	log.Warnw("warnw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	log.Errorw("errorw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
	// log.Fatalw("fatalw message", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))

	log.With(context.Background()).Infow("test with context", mo.Value("k1", 123), mo.Value("k2", true), mo.Value("k3", []int{1, 2, 3}))
}
