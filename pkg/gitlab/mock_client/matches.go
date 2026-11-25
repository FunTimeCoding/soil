package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_filter"
	"slices"
	"strings"
)

func matches(
	r *merge_request.Request,
	f *merge_request_filter.Filter,
) bool {
	switch {
	case r.Project != f.Project:
		return false
	case f.State != constant.AllState && r.State != f.State:
		return false
	case f.UpdatedAfter != nil &&
		(r.Update == nil || !r.Update.After(*f.UpdatedAfter)):

		return false
	case f.Author != "" && r.Author != f.Author:
		return false
	case f.Reviewer != "" && !slices.Contains(r.Reviewers, f.Reviewer):
		return false
	case f.Assignee != "" && !slices.Contains(r.Assignees, f.Assignee):
		return false
	case f.SourceBranch != "" && r.SourceBranch != f.SourceBranch:
		return false
	case f.TargetBranch != "" && r.TargetBranch != f.TargetBranch:
		return false
	case f.Search != "" && !strings.Contains(r.Title, f.Search):
		return false
	case f.Draft != nil && r.Draft != *f.Draft:
		return false
	default:
		return true
	}
}
