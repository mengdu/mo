package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/mengdu/mo"
)

type user struct {
	Name      string
	Email     string
	CreatedAt time.Time
}

var oneUser = &user{
	Name:      "Jane Doe",
	Email:     "jane@test.com",
	CreatedAt: time.Date(1980, 1, 1, 12, 0, 0, 0, time.UTC),
}
var meta = mo.Meta{
	"int":     1,
	"ints":    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	"string":  "hello",
	"strings": []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
	"time":    time.Unix(0, 0),
	"times": []time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(2, 0),
		time.Unix(3, 0),
		time.Unix(4, 0),
		time.Unix(5, 0),
		time.Unix(6, 0),
		time.Unix(7, 0),
		time.Unix(8, 0),
		time.Unix(9, 0),
	},
	"user1": oneUser,
	"users": []*user{
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
		oneUser,
	},
	"error": errors.New("fail"),
}

func main() {
	fmt.Println("Std Logger")
	mo.Error("Std Error message")
	mo.Warn("Std Warn message")
	mo.Info("Std Info message")
	mo.Log("Std Log message")
	mo.Success("Std Success message")
	mo.Debug("Std Debug message")
	mo.Errorf("Format message %d", 1)
	mo.Warnf("Format message %d", 1)
	mo.Infof("Format message %d", 1)
	mo.Logf("Format message %d", 1)
	mo.Successf("Format message %d", 1)
	mo.Debugf("Format message %d", 1)

	mo.With(map[string]interface{}{
		"a": 1,
	}).Info("Std With meta message")
	// <-time.Tick(time.Second)
	fmt.Println("Text formater logger")
	logger := mo.New()
	// file, _ := os.OpenFile("my.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// errfile, _ := os.OpenFile("my-err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// logger.Stdout = file
	// logger.Stderr = errfile
	logger.Level = mo.LEVEL_ALL
	// logger.Level = mo.WARN
	logger.Caller = true
	logger.RelativeFilePath = true
	logger.Meta = map[string]interface{}{
		"service": "demo",
	}
	// logger.Formater = &mo.JsonForamter{}
	logger.Formater = &mo.TextForamter{
		// DisableLevelIcon: true,
		EnableTime: true,
		TimeLayout: "15:04:05.000",
		// TimeLayout: "03:04:05.000PM",
		// DisableColor: true,
		EnableLevel: true,
		ShortLevel:  true,
	}

	logger.Error("Error message")
	logger.Warn("Warn message", 1, 2, 3)
	logger.Info("Info message")
	logger.Log("Log message")
	logger.Success("Success message")
	logger.Debug("Debug message")
	logger.Errorf("Format message %d", 1)
	logger.Warnf("Format message %d", 1)
	logger.Infof("Format message %d", 1)
	logger.Logf("Format message %d", 1)
	logger.Successf("Format message %d", 1)
	logger.Debugf("Format message %d", 1)

	l := logger.With(mo.Meta{
		"a":   "1",
		"arr": []string{"a", "b", "c"},
		// "log": logger,
	})
	l.Info("With meta message")
	logger.With(mo.Meta{
		"b": 1,
	}).With(mo.Meta{
		"c": 1,
	}).Info("With meta message")

	fmt.Println("JSON formater logger")
	jlog := mo.New()
	jlog.Formater = &mo.JsonForamter{}
	jlog.Meta = mo.Meta{
		"a": 1,
	}

	jlog.Error("Error message")
	jlog.Warn("Warn message", 1, 2, 3)
	jlog.Info("Info message")
	jlog.Log("Log message")
	jlog.Success("Success message")
	jlog.Debug("Debug message")
	jlog.Log("Print other")
	jlog.Errorf("Format message %d", 1)
	jlog.Warnf("Format message %d", 1)
	jlog.Infof("Format message %d", 1)
	jlog.Logf("Format message %d", 1)
	jlog.Successf("Format message %d", 1)
	jlog.Debugf("Format message %d", 1)
}
