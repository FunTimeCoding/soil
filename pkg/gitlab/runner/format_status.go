package runner

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func (r *Runner) formatStatus(f *option.Format) string {
	result := r.Status

	if f.UseColor {
		if result == constant.RunnerOnlineStatus {
			result = consoleConstant.Green("%s", result)
		} else {
			result = consoleConstant.Red("%s", result)
		}
	}

	return result
}
