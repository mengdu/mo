# Mo

![Preview](preview.png)

A leveled logger library for golang.

```sh
go get github.com/mengdu/mo
```

```go
package main
import (
	"github.com/mengdu/mo"
)

func main() {
	mo.SetBase(
		mo.Value("ts", mo.Timestamp("15:04:05.000")),
		mo.Value("caller", mo.Caller(4)),
		mo.Value("tag", "dev"),
	)

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
```

- [Example](examples/main.go)
- [Rotate Example](examples/rotate/main.go)
- [Opentelemetry Example](examples/otel/main.go)

## Benchmark

```txt
goos: darwin
goarch: amd64
pkg: github.com/mengdu/mo
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkDefault-4                      41029287               155.1 ns/op            64 B/op          2 allocs/op
BenchmarkDefaultWithCaller-4             9756843               550.4 ns/op           496 B/op         11 allocs/op
BenchmarkJson-4                         16316077               368.8 ns/op           128 B/op          6 allocs/op
BenchmarkJsonWithCaller-4                4433232              1470 ns/op             536 B/op         14 allocs/op
BenchmarkJsonWithCallerFull-4            2330289              2311 ns/op             816 B/op         21 allocs/op
PASS
ok      github.com/mengdu/mo    39.824s
```
