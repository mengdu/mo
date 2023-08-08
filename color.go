package mo

import (
	"fmt"
	"hash/fnv"
	"strings"
)

func color(s string, start string, end string) string {
	s = strings.ReplaceAll(s, "\u001b[39m", fmt.Sprintf("\u001b[39m\u001b[%sm", start))
	return fmt.Sprintf("\u001b[%sm%s\u001b[%sm", start, s, end)
}

func strHashCode(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	min := 1
	max := 231
	size := max - min + 1
	hashValue := h.Sum32()
	return hashValue%uint32(size) + uint32(min)
}
