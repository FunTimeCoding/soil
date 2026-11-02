package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
)

func (s *Service) Pipelines(
	instance string,
	project int64,
	reference string,
	status string,
	limit int64,
) ([]*pipeline.Pipeline, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*pipeline.Pipeline, error) {
			return c.Pipelines(project, reference, status, limit)
		},
	)
}
