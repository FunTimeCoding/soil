package test_run

import "github.com/funtimecoding/soil/pkg/test_event/constant"

func (r *Run) Passed() []string {
	return r.named(constant.ActionPass)
}
