package reporter

import "runtime"

func frames(skip int) []runtime.Frame {
	pcs := make([]uintptr, 64)
	n := runtime.Callers(skip+2, pcs)
	iterator := runtime.CallersFrames(pcs[:n])
	var result []runtime.Frame

	for {
		f, more := iterator.Next()

		if f.Function == "runtime.gopanic" {
			result = nil
		} else {
			result = append(result, f)
		}

		if !more {
			return result
		}
	}
}
