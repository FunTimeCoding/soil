package runner

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func (r *Runner) formatStatus(f *option.Format) string {
	result := r.Status

	if f.UseColor {
		if result == constant.RunnerOnlineStatus {
			result = console.Green("%s", result)
		} else {
			result = console.Red("%s", result)
		}
	}

	return result
}
