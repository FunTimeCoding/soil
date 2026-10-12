package job

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	library "github.com/funtimecoding/soil/pkg/time"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func (j *Job) formatDate(f *option.Format) string {
	var format string
	t := j.Create.Local()

	if f.HasTag(consoleConstant.TagDense) && t.After(library.Midnight(time.Now())) {
		format = timeConstant.HourMinute
	} else {
		format = timeConstant.DateMinute
	}

	return t.Format(format)
}
