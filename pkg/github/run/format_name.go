package run

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Run) formatName(f *option.Format) string {
	result := r.Repository.FullName

	if f.UseColor {
		result = consoleConstant.Cyan("%s", result)
	}

	return result
}
