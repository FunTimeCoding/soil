package runner

import "github.com/funtimecoding/soil/pkg/gitlab/constant"

func (r *Runner) formatDescription() string {
	if r.Description == "" {
		return constant.RunnerNoDescription
	}

	return r.Description
}
