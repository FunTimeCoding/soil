package job

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (j *Job) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", j.Name)
	}

	return j.Name
}
