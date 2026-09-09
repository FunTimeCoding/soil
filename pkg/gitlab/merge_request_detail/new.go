package merge_request_detail

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.MergeRequest) *Detail {
	return &Detail{
		Project:    v.ProjectID,
		Identifier: v.IID,
		Title:      v.Title,
		State:      v.State,
		Link:       v.WebURL,
		Create:     v.CreatedAt,
		Raw:        v,
	}
}
