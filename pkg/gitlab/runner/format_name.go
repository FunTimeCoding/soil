package runner

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func (r *Runner) formatName(f *option.Format) string {
	result := r.Name

	if result == "" {
		result = constant.RunnerNoName
	}

	if f.UseColor {
		result = consoleConstant.Cyan("%s", result)
	}

	return result
}
