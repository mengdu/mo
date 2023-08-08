package main

import (
	"fmt"

	"github.com/mengdu/mo"
)

func main() {
	fmt.Println("Std Logger")
	mo.Std.Tag = "A"
	mo.Error("Std Error message")
	mo.Warn("Std Warn message")
	mo.Info("Std Info message")
	mo.Log("Std Log message")
	mo.Success("Std Success message")
	mo.Debug("Std Debug message")
	mo.Std.Tag = "B"
	mo.Errorf("Format message %ds", 1)
	mo.Warnf("Format message %ds", 1)
	mo.Infof("Format message %ds", 1)
	mo.Logf("Format message %ds", 1)
	mo.Successf("Format message %ds", 1)
	mo.Debugf("Format message %ds", 1)
	mo.Warnf("Bool %t, Int %d, Float %f, String %s", true, 666, 3.24, "Hello")
	// mo.Panic("Something Wrong!")
	// mo.Panicf("%d Wrongs!", 1)

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
	// logger.DisableColor = true
	logger.Level = mo.LEVEL_ALL
	// logger.Level = mo.WARN
	logger.Tag = "HTTP"
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
		EnableLevel: true,
		ShortLevel:  true,
	}

	logger.Error("Error message")
	// logger.Panic("Panic message")
	logger.Warn("Warn message", 1, 2, 3)
	logger.Info("Info message")
	logger.Log("Log message")
	logger.Success("Success message")
	logger.Debug("Debug message")
	logger.Errorf("Format message %d", 1)
	// logger.Panicf("Format message %d", 1)
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
	}).Infof("With meta message %s", "Hello")

	fmt.Println("JSON formater logger")
	jlog := mo.New()
	jlog.Tag = "HTTP"
	jlog.DisableSprintfColor = true
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

	fmt.Println(logger.Sprintf("var: %v\nvar+: %+v\nvar#: %#v\n", logger, logger, logger))
	fmt.Println(logger.Sprintf("T: %T\nt: %t\np: %p\n", logger, true, logger))
	fmt.Println(logger.Sprintf("b: %b\no: %o\nx: %x\nX: %X\nd: %d\n", 8273412, 8273412, 8273412, 8273412, 8273412))
	fmt.Println(logger.Sprintf("s: %s\nU: %U\nq: %q", "Hello Mo!", '中', "Hi, \nMo!"))
}
