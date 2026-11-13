package test_run

import (
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/test_event/constant"
)

func (r *Run) BuildOutput() []string {
	var result []string

	for _, e := range r.events {
		if e.Action == constant.ActionBuildOutput {
			result = append(
				result,
				slice.StripEmpty(split.NewLine(e.Output))...)
		}
	}

	return result
}
