package mo

import "runtime"

func getCaller(skip int) runtime.Frame {
	pcs := make([]uintptr, 1)
	n := runtime.Callers(skip, pcs)
	frames := runtime.CallersFrames(pcs[0:n])
	frame, _ := frames.Next()
	return frame
}
