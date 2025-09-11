package mo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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
	cwd     string
	cwdOnce sync.Once
)

type ICaller interface {
	Caller(skip int) string
}

// static check
var _ ICaller = caller{}

type caller struct{}

func (caller) Caller(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	cwdOnce.Do(func() {
		cwd, _ = os.Getwd()
	})
	if strings.Index(file, cwd) == 0 {
		file, _ = filepath.Rel(cwd, file)
	}
	return file + ":" + strconv.Itoa(line)
}

func Caller() Valuer {
	return func(ctx context.Context) interface{} {
		return caller{}
	}
}

func Timestamp(layout string) Valuer {
	return func(context.Context) interface{} {
		return time.Now().Format(layout)
	}
}
