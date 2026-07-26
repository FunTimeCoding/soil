package model_context

import (
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func formatTime(t time.Time) string {
	return t.Format(constant.DateMinute)
}
