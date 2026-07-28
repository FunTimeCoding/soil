package run

import (
	"github.com/funtimecoding/soil/pkg/github/constant"
	"slices"
)

func (r *Run) Validate() {
	if len(r.Jobs) > 0 && r.Jobs[0].Fail() {
		if !slices.Contains(r.concern, constant.RunFailedConcern) {
			r.concern = append(r.concern, constant.RunFailedConcern)
		}
	}
}
