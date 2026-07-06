package merge_request

import "github.com/funtimecoding/soil/pkg/gitlab/constant"

func (r *Request) Done() bool {
	if r.State == constant.MergedState {
		return true
	}

	if r.State == constant.ClosedState {
		return true
	}

	return false
}
