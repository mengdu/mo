module github.com/mengdu/mo/record

go 1.23.0

toolchain go1.24.10

replace "github.com/mengdu/mo" => ../

require (
	github.com/mengdu/color v0.4.0
	github.com/mengdu/mo v1.3.1
	github.com/natefinch/lumberjack v2.0.0+incompatible
	go.opentelemetry.io/otel/trace v1.38.0
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	go.opentelemetry.io/otel v1.38.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
