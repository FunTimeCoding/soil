package helper

import (
	"github.com/funtimecoding/soil/pkg/linux/constant"
	"time"
)

func ParseTimestamp(s string) time.Time {
	result, e := time.Parse(constant.SystemdDateTime, s)

	if e != nil {
		return time.Time{}
	}

	return result
}
