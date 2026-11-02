package loki

import timeConstant "github.com/funtimecoding/soil/pkg/time/constant"

func formatLatest(e *Overview) string {
	if e.Latest.IsZero() {
		return ""
	}

	return e.Latest.Format(timeConstant.DateMinute)
}
