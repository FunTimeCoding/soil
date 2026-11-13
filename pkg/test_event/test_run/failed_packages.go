package test_run

import "github.com/funtimecoding/soil/pkg/test_event/constant"

func (r *Run) FailedPackages() []string {
	var result []string

	for _, e := range r.events {
		if e.Action == constant.ActionFail && e.Test == "" {
			result = append(result, e.Package)
		}
	}

	return sorted(result)
}
