package pipeline_detail

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.Pipeline) *Detail {
	return &Detail{
		Identifier:        v.ID,
		ProjectIdentifier: v.ProjectID,
		Status:            v.Status,
		Source:            string(v.Source),
		Reference:         v.Ref,
		Hash:              v.SHA,
		Link:              v.WebURL,
		Create:            v.CreatedAt,
		Update:            v.UpdatedAt,
		Raw:               v,
	}
}
