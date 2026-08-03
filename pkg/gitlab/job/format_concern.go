package job

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (j *Job) formatConcern(f *option.Format) string {
	if len(j.concern) == 0 {
		return ""
	}

	result := join.Comma(j.concern)

	if f.UseColor {
		result = constant.Yellow("%s", result)
	}

	return result
}
