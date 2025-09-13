package mo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Field [2]interface{}

func (v Field) Key() string {
	if v, ok := v[0].(string); ok {
		return v
	}
	return fmt.Sprint(v[0])
}

func (v Field) Value() interface{} {
	return v[1]
}

func Value(key string, value interface{}) Field {
	return Field{key, value}
}

type Valuer func(ctx context.Context) interface{}

var (
	cwd       string
	cwdOnce   sync.Once
	callerAbs = os.Getenv("MO_CALLER_ABS") == "1"
)

func Caller(skip int) Valuer {
	return func(ctx context.Context) interface{} {
		_, file, line, _ := runtime.Caller(skip)
		cwdOnce.Do(func() {
			cwd, _ = os.Getwd()
		})

		if runtime.GOOS == "windows" {
			file = filepath.Clean(file)
		}

		if !callerAbs && strings.Index(file, cwd) == 0 {
			file, _ = filepath.Rel(cwd, file)
		}
		return fmt.Sprintf("%s:%d", file, line)
	}
}

func Timestamp(layout string) Valuer {
	return func(context.Context) interface{} {
		return time.Now().Format(layout)
	}
}
