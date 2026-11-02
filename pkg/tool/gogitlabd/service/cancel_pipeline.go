package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
)

func (s *Service) CancelPipeline(
	instance string,
	project int64,
	identifier int64,
) (*pipeline_detail.Detail, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*pipeline_detail.Detail, error) {
			return c.CancelPipeline(project, identifier)
		},
	)
}
