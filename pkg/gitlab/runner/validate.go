package runner

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"slices"
)

func (r *Runner) Validate() {
	if r.Paused && !slices.Contains(r.concern, constant.RunnerPaused) {
		r.concern = append(r.concern, constant.RunnerPaused)
	}

	if r.Address == "" && !slices.Contains(
		r.concern,
		constant.RunnerNoAddressConcern,
	) {
		r.concern = append(r.concern, constant.RunnerNoAddressConcern)
	}
}
