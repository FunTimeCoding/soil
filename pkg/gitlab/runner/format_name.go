package runner

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func (r *Runner) formatName(f *option.Format) string {
	result := r.Name

	if result == "" {
		result = constant.RunnerNoName
	}

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
