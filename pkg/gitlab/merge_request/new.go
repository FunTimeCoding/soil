package merge_request

import "gitlab.com/gitlab-org/api/client-go/v3"

func New(v *gitlab.BasicMergeRequest) *Request {
	result := &Request{
		Project:      v.ProjectID,
		Identifier:   v.IID,
		Title:        v.Title,
		State:        v.State,
		Link:         v.WebURL,
		Create:       v.CreatedAt,
		Update:       v.UpdatedAt,
		Reviewers:    usernames(v.Reviewers),
		Assignees:    usernames(v.Assignees),
		SourceBranch: v.SourceBranch,
		TargetBranch: v.TargetBranch,
		Draft:        v.Draft,
		MergeStatus:  v.DetailedMergeStatus,
		Notes:        v.UserNotesCount,
		Raw:          v,
	}

	if v.Author != nil {
		result.Author = v.Author.Username
	}

	return result
}
