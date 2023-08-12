package buffer

import (
	"bytes"
	"sync"
)

var pool *sync.Pool

func init() {
	pool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

func Get() *bytes.Buffer {
	return pool.Get().(*bytes.Buffer)
}

func Put(b *bytes.Buffer) {
	b.Reset()
	pool.Put(b)
}

func Append(b *bytes.Buffer, value string) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(value)
}
