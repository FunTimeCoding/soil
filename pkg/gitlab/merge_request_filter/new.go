package merge_request_filter

import "github.com/funtimecoding/soil/pkg/gitlab/constant"

func New(project int64) *Filter {
	return &Filter{
		Project: project,
		Limit:   constant.RequestLimit,
		State:   constant.OpenedState,
		OrderBy: constant.UpdatedOrder,
	}
}
