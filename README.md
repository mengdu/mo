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
go test -benchmem -bench "^Benchmark" -benchtime=5s
goos: darwin
goarch: amd64
pkg: github.com/mengdu/mo
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_Info-12                       30961867               186.8 ns/op            16 B/op          1 allocs/op
Benchmark_Infof-12                      25124493               202.2 ns/op            24 B/op          1 allocs/op
Benchmark_Infow-12                      16628475               423.4 ns/op           272 B/op          8 allocs/op
Benchmark_Infox-12                      27565543               205.0 ns/op            16 B/op          1 allocs/op
Benchmark_Infofx-12                     28369299               219.1 ns/op            24 B/op          1 allocs/op
Benchmark_Infowx-12                     14071964               462.8 ns/op           288 B/op          9 allocs/op
Benchmark_WithCaller_Info-12            13652521               438.9 ns/op           457 B/op         10 allocs/op
Benchmark_WithCaller_Infof-12           12920677               479.2 ns/op           465 B/op         10 allocs/op
Benchmark_WithCaller_Infow-12            6882974               898.2 ns/op           714 B/op         16 allocs/op
Benchmark_WithCaller_Infox-12           11326057               515.0 ns/op           457 B/op         10 allocs/op
Benchmark_WithCaller_Infofx-12          10696102               499.6 ns/op           465 B/op         10 allocs/op
Benchmark_WithCaller_Infowx-12           6628028              1002 ns/op             730 B/op         17 allocs/op
Benchmark_With-12                       10844205               519.7 ns/op           457 B/op         10 allocs/op
Benchmark_JSON-12                        1522146              3488 ns/op            1585 B/op         34 allocs/op
PASS
ok      github.com/mengdu/mo    98.779s
```
