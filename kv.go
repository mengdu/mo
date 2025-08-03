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

type KeyValue [2]interface{}

func (v KeyValue) Key() string {
	if v, ok := v[0].(string); ok {
		return v
	}
	return fmt.Sprint(v[0])
}

func (v KeyValue) Value() interface{} {
	return v[1]
}

func Value(key string, value interface{}) KeyValue {
	return KeyValue{key, value}
}

type Valuer func(ctx context.Context) interface{}

var (
	cwd     string
	cwdOnce sync.Once
)

func Caller(depth int) Valuer {
	return func(ctx context.Context) interface{} {
		_, file, line, _ := runtime.Caller(depth)
		cwdOnce.Do(func() {
			cwd, _ = os.Getwd()
		})
		if strings.Index(file, cwd) == 0 {
			file, _ = filepath.Rel(cwd, file)
		}
		return file + ":" + strconv.Itoa(line)
	}
}

func Timestamp(layout string) Valuer {
	return func(context.Context) interface{} {
		return time.Now().Format(layout)
	}
}
