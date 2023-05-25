package mo

import (
	"bytes"
	"sync"
)

var pool sync.Pool

func newBuffer() *bytes.Buffer {
	if buf, ok := pool.Get().(*bytes.Buffer); ok {
		return buf
	}
	return &bytes.Buffer{}
}

func releaseBuffer(buf *bytes.Buffer) {
	buf.Reset()
	pool.Put(buf)
}

func appendValue(b *bytes.Buffer, value string) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(value)
}
