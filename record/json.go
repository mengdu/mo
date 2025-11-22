package record

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/mengdu/mo"
)

// Key constants for structured logging
type JSON struct {
	Encoder *json.Encoder
	mu      sync.Mutex
}

// Log implements the Recorder interface.
func (l *JSON) Log(ctx context.Context, level mo.Level, msg string, kv []mo.Field) {
	line := make(map[string]interface{}, len(kv)+2)
	line[KeyLevel] = strings.ToLower(level.String())
	line[KeyMessage] = msg

	for _, v := range kv {
		line[v.Key()] = v.Value()
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if err := l.Encoder.Encode(line); err != nil {
		fmt.Fprintf(os.Stderr, "write failed: %v\n", err)
	}
}
