package job

import (
	console "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func (j *Job) formatDate(f *option.Format) string {
	var format string
	t := j.Create.Local()

	if f.HasTag(console.TagDense) && t.After(library.Midnight(time.Now())) {
		format = constant.HourMinute
	} else {
		format = constant.DateMinute
	}

	return t.Format(format)
}
