module oteldemo

go 1.23.0

replace github.com/mengdu/mo => ../../

require (
	github.com/mengdu/color v0.4.0
	github.com/mengdu/fmtx v0.3.1
	github.com/mengdu/mo v0.5.1
	go.opentelemetry.io/otel/trace v1.37.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/mattn/go-isatty v0.0.20 // indirect
	go.opentelemetry.io/otel v1.37.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
)
