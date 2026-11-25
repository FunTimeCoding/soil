package change

import (
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
)

func Actor(d *merge_request_detail.Detail) string {
	v := d.Raw

	switch {
	case v.State == gitlabConstant.MergedState && v.MergeUser != nil:
		return v.MergeUser.Username
	case v.State == gitlabConstant.ClosedState && v.ClosedBy != nil:
		return v.ClosedBy.Username
	default:
		return ""
	}
}
