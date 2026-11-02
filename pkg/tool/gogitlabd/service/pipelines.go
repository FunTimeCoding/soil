package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
)

func (s *Service) Pipelines(
	x context.Context,
	instance string,
	project int64,
	reference string,
	status string,
	limit int64,
) ([]*pipeline.Pipeline, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*pipeline.Pipeline, error) {
			return c.Pipelines(y, project, reference, status, limit)
		},
	)
}
