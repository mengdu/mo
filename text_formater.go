package mo

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mengdu/mo/buffer"
)

func levelIcon(l Level) string {
	switch l {
	case LEVEL_ERROR:
		// return "✗"
		// return "✕"
		return "✘"
	case LEVEL_WARN:
		// return "⚠"
		return "✧"
		// return "✦"
	case LEVEL_INFO:
		// return "і"
		return "⬗"
		// return "◎"
	case LEVEL_LOG:
		// return "⌾"
		// return "๏"
		return "○"
	case LEVEL_SUCCESS:
		return "✔"
	case LEVEL_DEBUG:
		// return "ⱗ"
		// return "ӿ"
		// return "⚜"
		return "❅"
	}
	return ""
}

func levelShort(l Level) string {
	switch l {
	case LEVEL_ERROR:
		return "ERR"
	case LEVEL_WARN:
		return "WRN"
	case LEVEL_INFO:
		return "INF"
	case LEVEL_LOG:
		return "LOG"
	case LEVEL_SUCCESS:
		return "SUC"
	case LEVEL_DEBUG:
		return "DBG"
	}
	return ""
}

type TextForamter struct {
	DisableLevelIcon bool
	EnableTime       bool
	EnableLevel      bool
	ShortLevel       bool
	TimeLayout       string
}

func (f *TextForamter) Format(log *Record) ([]byte, error) {
	icon := ""
	msg := log.Message
	file := log.Filename
	meta := ""
	at := ""
	tag := log.Tag
	level := fmt.Sprintf("[%s]", strings.ToUpper(log.Level.String()))

	if !f.DisableLevelIcon {
		icon = levelIcon(log.Level)
	}

	if f.ShortLevel {
		level = fmt.Sprintf("[%s]", levelShort(log.Level))
	}
	if f.EnableTime {
		layout := f.TimeLayout
		if layout == "" {
			layout = "2006-01-02T15:04:05.000Z07:00"
		}
		at = fmt.Sprintf("[%s]", log.At.Format(layout))
	}

	if len(log.Meta) > 0 {
		buf, err := json.Marshal(log.Meta)
		if err != nil {
			return []byte(""), err
		}
		meta = string(buf)
	}

	if log.Logger.EnableColor() {
		if tag != "" {
			tag = color(fmt.Sprintf("[%s]", log.Tag), fmt.Sprintf("38;5;%d", strHashCode(log.Tag)), "39")
		}

		if icon != "" {
			switch log.Level {
			case LEVEL_ERROR:
				icon = color(icon, "31", "39")
			case LEVEL_WARN:
				icon = color(icon, "93", "39")
			case LEVEL_INFO:
				icon = color(icon, "36", "39")
			case LEVEL_LOG:
				icon = color(icon, "34", "39")
			case LEVEL_SUCCESS:
				icon = color(icon, "32", "39")
			case LEVEL_DEBUG:
				icon = color(icon, "33", "39")
			}
		}

		if f.EnableLevel {
			switch log.Level {
			case LEVEL_ERROR:
				level = color(level, "31", "39")
				msg = color(msg, "31", "39")
			case LEVEL_WARN:
				level = color(level, "93", "39")
				msg = color(msg, "93", "39")
			case LEVEL_INFO:
				level = color(level, "36", "39")
			case LEVEL_LOG:
				level = color(level, "34", "39")
			case LEVEL_SUCCESS:
				level = color(level, "32", "39")
			case LEVEL_DEBUG:
				level = color(level, "33", "39")
			}
		}

		switch log.Level {
		case LEVEL_ERROR:
			msg = color(msg, "31", "39")
		case LEVEL_WARN:
			msg = color(msg, "93", "39")
		}

		if file != "" {
			file = color(file, "2", "22;0;39")
		}
		if meta != "" {
			meta = color(meta, "34", "39")
		}
		if at != "" {
			at = color(at, "2;37", "0;39")
		}
	} else {
		if tag != "" {
			tag = fmt.Sprintf("[%s]", tag)
		}
	}

	buf := log.Buf
	spaceRequired := false
	if icon != "" {
		buffer.Append(buf, icon)
	}
	if at != "" {
		buffer.Append(buf, at)
		spaceRequired = true
	}
	if tag != "" {
		if !spaceRequired {
			buffer.Append(buf, tag)
		} else {
			buf.Write([]byte(tag))
		}
		spaceRequired = true
	}
	if f.EnableLevel {
		if !spaceRequired {
			buffer.Append(buf, level)
		} else {
			buf.Write([]byte(level))
		}
		spaceRequired = true
	}
	buffer.Append(buf, msg)
	if meta != "" {
		buffer.Append(buf, meta)
	}
	if file != "" {
		buffer.Append(buf, file)
	}
	buf.WriteByte('\n')
	return buf.Bytes(), nil
}
