package merge_request

import "gitlab.com/gitlab-org/api/client-go/v3"

func New(v *gitlab.BasicMergeRequest) *Request {
	return &Request{
		Project:    v.ProjectID,
		Identifier: v.IID,
		Title:      v.Title,
		State:      v.State,
		Link:       v.WebURL,
		Create:     v.CreatedAt,
		Raw:        v,
	}
}
