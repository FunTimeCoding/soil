package merge_request_filter

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (f *Filter) Options(page int64) *gitlab.ListProjectMergeRequestsOptions {
	result := &gitlab.ListProjectMergeRequestsOptions{
		State:        new(f.State),
		OrderBy:      new(f.OrderBy),
		Sort:         new(constant.Descending),
		UpdatedAfter: f.UpdatedAfter,
		Draft:        f.Draft,
		ListOptions: gitlab.ListOptions{
			Page:    page,
			PerPage: int64(f.Limit),
		},
	}

	if f.Author != "" {
		result.AuthorUsername = new(f.Author)
	}

	if f.Reviewer != "" {
		result.ReviewerUsername = new(f.Reviewer)
	}

	if f.Assignee != "" {
		result.AssigneeUsername = &[]string{f.Assignee}
	}

	if f.SourceBranch != "" {
		result.SourceBranch = new(f.SourceBranch)
	}

	if f.TargetBranch != "" {
		result.TargetBranch = new(f.TargetBranch)
	}

	if f.Search != "" {
		result.Search = new(f.Search)
	}

	return result
}
