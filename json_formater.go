package mo

import (
	"encoding/json"

	"github.com/mengdu/mo/buffer"
)

type JsonForamter struct{}

func (f *JsonForamter) Format(log *Record) ([]byte, error) {
	buf := buffer.Get()
	defer buffer.Put(buf)
	err := json.NewEncoder(buf).Encode(log)
	return buf.Bytes(), err
}
