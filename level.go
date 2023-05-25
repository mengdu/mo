package mo

import "fmt"

type Level int8

func (l Level) String() string {
	if b, err := l.MarshalText(); err == nil {
		return string(b)
	} else {
		return "unknown"
	}
}

func (l Level) MarshalText() ([]byte, error) {
	switch l {
	case LEVEL_NONE:
		return []byte("none"), nil
	case LEVEL_ERROR:
		return []byte("error"), nil
	case LEVEL_WARN:
		return []byte("warn"), nil
	case LEVEL_INFO:
		return []byte("info"), nil
	case LEVEL_LOG:
		return []byte("log"), nil
	case LEVEL_SUCCESS:
		return []byte("success"), nil
	case LEVEL_DEBUG:
		return []byte("debug"), nil
	case LEVEL_ALL:
		return []byte("all"), nil
	}
	return []byte(fmt.Sprintf("%d", l)), nil
}

const (
	LEVEL_NONE Level = iota
	LEVEL_ERROR
	LEVEL_WARN
	LEVEL_INFO
	LEVEL_LOG
	LEVEL_SUCCESS
	LEVEL_DEBUG
	LEVEL_ALL
)

func isEnableLevel(a Level, b Level) bool {
	return b <= a
}
