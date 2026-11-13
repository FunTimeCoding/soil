package test_run

import (
	"github.com/funtimecoding/soil/pkg/test_event/constant"
	"strings"
)

func (r *Run) Panicked() []string {
	var result []string

	for _, e := range r.events {
		if e.Action == constant.ActionOutput &&
			e.Test != "" &&
			strings.HasPrefix(e.Output, constant.PanicPrefix) {
			result = append(result, e.Test)
		}
	}

	return sorted(result)
}
