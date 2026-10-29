package runtime

import "runtime"

func CallerName(skip int) string {
	program := make([]uintptr, 1)

	if runtime.Callers(skip+2, program) == 0 {
		return ""
	}

	frame, _ := runtime.CallersFrames(program).Next()

	return frame.Function
}
