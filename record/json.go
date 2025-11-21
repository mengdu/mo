package record

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mengdu/mo"
)

type JSONRecorder struct {
	Encoder *json.Encoder
}

func (l *JSONRecorder) Log(ctx context.Context, level mo.Level, msg string, kv []mo.Field) {
	line := make(map[string]interface{}, len(kv)+2)
	line["level"] = strings.ToLower(level.String())
	line["msg"] = msg

	for _, v := range kv {
		line[v.Key()] = v.Value()
	}

	if err := l.Encoder.Encode(line); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v", err))
	}
}
