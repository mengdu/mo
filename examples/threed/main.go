package main

import (
	"encoding/json"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/mengdu/mo"
	"github.com/mengdu/mo/buffer"
)

type Liner struct {
	raw     []byte
	lineCnt int32
	mu      sync.Mutex
}

func (l *Liner) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	atomic.AddInt32(&l.lineCnt, 1)
	l.raw = append(l.raw, p...)
	return len(p), nil
}

type MyForamter struct {
	mu sync.Mutex
}

func (f *MyForamter) Format(log *mo.Record) ([]byte, error) {
	// // fail
	// f.mu.Lock()
	// defer f.mu.Unlock()
	buf := buffer.Get()
	defer buffer.Put(buf)
	err := json.NewEncoder(buf).Encode(log)
	return buf.Bytes(), err

	// // ok
	// buf := buffer.Get()
	// defer buffer.Put(buf)
	// buf.WriteString(log.Message + "\n")
	// return buf.Bytes(), nil

	// // ok
	// buf := &bytes.Buffer{}
	// err := json.NewEncoder(buf).Encode(log)
	// return buf.Bytes(), err

	// // ok
	// buf, err := json.Marshal(log)
	// buf = append(buf, '\n')
	// return buf, err
}

func ThreedTest() {
	liner := &Liner{}
	log := mo.New()
	log.Stderr = liner
	log.Stdout = liner
	log.Caller = true
	log.RelativeFilePath = true
	// log.DisableColor = true
	// log.Formater = &mo.JsonForamter{}
	log.Formater = &MyForamter{}
	w := sync.WaitGroup{}
	i := int32(0)
	n := int32(2)
	// demo := Demo{
	// 	writer:   liner,
	// 	formater: &DemoFormater{},
	// }
	for ; i < 1000; i++ {
		w.Add(int(n))
		go func() {
			// demo.Log("line message 1")
			log.Warn("line message 1")
			// log.WithTag("A").Warn("line message 1")
			// liner.Write([]byte("line message 1\n"))
			w.Done()
		}()
		go func() {
			// demo.Log("line message 2")
			log.Info("line message 2")
			// log.WithTag("B").Info("line message 2")
			// liner.Write([]byte("line message 2\n"))
			w.Done()
		}()
	}
	w.Wait()
	cnt := i * n
	lines := strings.Split(strings.TrimSpace(string(liner.raw)), "\n")
	length := int32(len(lines))
	// fmt.Print(string(liner.raw))
	if liner.lineCnt != cnt {
		mo.Errorf("Invoke Expect %d, Got %d", cnt, liner.lineCnt)
	} else if length != liner.lineCnt {
		// fmt.Print(string(liner.raw))
		mo.Errorf("Log Line Expect %d, Got %d", liner.lineCnt, length)
	} else {
		// mo.Successf("Test ok!")
	}
}

func main() {
	for i := 0; i < 25; i++ {
		ThreedTest()
	}
}
