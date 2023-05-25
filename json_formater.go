package mo

import (
	"encoding/json"
)

type JsonForamter struct{}

func (f *JsonForamter) Format(log *Record) ([]byte, error) {
	buf := newBuffer()
	defer releaseBuffer(buf)
	err := json.NewEncoder(buf).Encode(log)
	return buf.Bytes(), err
}
