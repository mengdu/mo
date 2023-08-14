package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/mengdu/mo"
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
	// mu sync.Mutex
}

func (f *MyForamter) Format(log *mo.Record) ([]byte, error) {
	err := json.NewEncoder(log.Buf).Encode(log)
	return log.Buf.Bytes(), err
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
	for ; i < 10000; i++ {
		w.Add(int(n))
		go func() {
			// log.Warn("line message 1")
			log.WithTag("A").Warnf("line message %d", 1)
			// liner.Write([]byte("line message 1\n"))
			w.Done()
		}()
		go func() {
			// log.Info("line message 2")
			log.WithTag("B").Infof("line message %d", 2)
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
		mo.Panicf("Invoke Expect %d, Got %d", cnt, liner.lineCnt)
	} else if length != liner.lineCnt {
		fmt.Print(string(liner.raw))
		mo.Panicf("Log Line Expect %d, Got %d", liner.lineCnt, length)
	} else {
		// mo.Successf("Test ok!")
	}
}

func main() {
	for i := 0; i < 25; i++ {
		ThreedTest()
	}
}
