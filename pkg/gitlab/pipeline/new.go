package pipeline

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.PipelineInfo) *Pipeline {
	return &Pipeline{
		Identifier:        v.ID,
		ProjectIdentifier: v.ProjectID,
		Status:            v.Status,
		Source:            v.Source,
		Reference:         v.Ref,
		Hash:              v.SHA,
		Link:              v.WebURL,
		Create:            v.CreatedAt,
		Update:            v.UpdatedAt,
	}
}
