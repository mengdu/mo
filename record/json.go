package record

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mengdu/mo"
)

// Key constants for structured logging
type JSON struct {
	Encoder *json.Encoder
}

// Log implements the Recorder interface.
func (l *JSON) Log(ctx context.Context, level mo.Level, msg string, kv []mo.Field) {
	line := make(map[string]interface{}, len(kv)+2)
	line[KeyLevel] = strings.ToLower(level.String())
	line[KeyMessage] = msg

	for _, v := range kv {
		line[v.Key()] = v.Value()
	}

	if err := l.Encoder.Encode(line); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v", err))
	}
}
