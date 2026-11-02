package job

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (j *Job) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", j.Name)
	}

	return j.Name
}
