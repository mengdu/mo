package mo

import (
	"encoding/json"
)

type JsonForamter struct{}

func (f *JsonForamter) Format(log *Record) ([]byte, error) {
	err := json.NewEncoder(log.Buf).Encode(log)
	return log.Buf.Bytes(), err
}
