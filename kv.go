package mo

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
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

var callerAbs, _ = strconv.ParseBool(os.Getenv("MO_CALLER_ABS"))

func Caller(skip int) Valuer {
	return func(ctx context.Context) interface{} {
		_, file, line, _ := runtime.Caller(skip)
		if callerAbs {
			return file + ":" + strconv.Itoa(line)
		}
		idx := strings.LastIndexByte(file, '/')
		if idx == -1 {
			return file[idx+1:] + ":" + strconv.Itoa(line)
		}
		idx = strings.LastIndexByte(file[:idx], '/')
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

func Timestamp(layout string) Valuer {
	return func(context.Context) interface{} {
		return time.Now().Format(layout)
	}
}
