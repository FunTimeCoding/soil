package minute

import (
	"github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/time/second"
)

func Readable(minutes int) string {
	return second.Readable(minutes * constant.MinuteInSeconds)
}
