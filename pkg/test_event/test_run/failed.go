package test_run

import "github.com/funtimecoding/soil/pkg/test_event/constant"

func (r *Run) Failed() []string {
	return r.named(constant.ActionFail)
}
