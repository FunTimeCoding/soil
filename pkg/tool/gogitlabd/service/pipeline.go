package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
)

func (s *Service) Pipeline(
	x context.Context,
	instance string,
	project int64,
	identifier int64,
) (*pipeline_detail.Detail, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*pipeline_detail.Detail, error) {
			return c.Pipeline(y, project, identifier)
		},
	)
}
