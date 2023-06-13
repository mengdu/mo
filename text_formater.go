package mo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

func color(s string, start string, end string) string {
	return fmt.Sprintf("\u001b[%sm%s\u001b[%sm", start, s, end)
}

func levelIcon(l Level) string {
	switch l {
	case LEVEL_ERROR:
		return "✗"
	case LEVEL_WARN:
		return "⚠"
		// return "✧"
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
	DisableColor     bool
	ForceColor       bool
	DisableLevelIcon bool
	EnableTime       bool
	EnableLevel      bool
	ShortLevel       bool
	TimeLayout       string
}

func (f *TextForamter) isColored() bool {
	isColored := f.ForceColor || isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
	switch force, ok := os.LookupEnv("FORCE_COLOR"); {
	case ok && force != "0":
		isColored = true
	}
	return isColored && !f.DisableColor
}

func (f *TextForamter) Format(log *Record) ([]byte, error) {
	icon := ""
	msg := log.Message
	file := log.Filename
	meta := ""
	at := ""
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

	if f.isColored() {
		if icon != "" {
			switch log.Level {
			case LEVEL_ERROR:
				icon = color(icon, "31", "0")
				level = color(level, "31", "0")
				msg = color(msg, "31", "0")
			case LEVEL_WARN:
				icon = color(icon, "93", "0")
				level = color(level, "93", "0")
				msg = color(msg, "93", "0")
			case LEVEL_INFO:
				icon = color(icon, "36", "0")
				level = color(level, "36", "0")
			case LEVEL_LOG:
				icon = color(icon, "34", "0")
				level = color(level, "34", "0")
			case LEVEL_SUCCESS:
				icon = color(icon, "32", "0")
				level = color(level, "32", "0")
			case LEVEL_DEBUG:
				icon = color(icon, "33", "0")
				level = color(level, "33", "0")
			}
		}
		if file != "" {
			file = color(file, "2", "22")
		}
		if log.Level == LEVEL_ERROR {
			msg = color(msg, "1;31", "0")
		}
		if meta != "" {
			meta = color(meta, "34", "0")
		}
		if at != "" {
			at = color(at, "2;37", "0")
		}
	}

	buf := newBuffer()
	defer releaseBuffer(buf)
	if icon != "" {
		appendValue(buf, icon)
	}
	if at != "" {
		appendValue(buf, at)
	}
	if f.EnableLevel {
		appendValue(buf, level)
	}
	appendValue(buf, msg)
	if meta != "" {
		appendValue(buf, meta)
	}
	if file != "" {
		appendValue(buf, file)
	}
	buf.WriteByte('\n')
	return buf.Bytes(), nil
}
