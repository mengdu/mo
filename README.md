# Mo

A leveled logger library for golang.

```sh
go get github.com/mengdu/mo
```

```go
package main
import (
	"github.com/mengdu/mo"
)

func main() {
	mo.Error("Error message")
	mo.Warn("Warn message")
	mo.Info("Warn message")
	mo.Log("Log message")
	mo.Success("Success message")
	mo.Debug("Debug message")
	mo.With(map[string]interface{}{
		"a": 1,
	}).Info("With meta message")
}
```

![Preview](preview.png)

## Levels

````
LEVEL_NONE
LEVEL_ERROR
LEVEL_WARN
LEVEL_INFO
LEVEL_LOG
LEVEL_SUCCESS
LEVEL_DEBUG
LEVEL_ALL
````

```go
logger := mo.New()
logger.Level = mo.LEVEL_ALL
```

## Production recommendations

```go
file, _ := os.OpenFile("service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
errfile, _ := os.OpenFile("service-err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
logger := mo.New()
logger.Formater = &mo.JsonForamter{}
logger.Stdout = file
logger.Stderr = errfile
logger.Level = mo.LEVEL_SUCCESS
```
