package time

import (
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func Format(t time.Time) string {
	return t.Format(constant.DateMinute)
}
