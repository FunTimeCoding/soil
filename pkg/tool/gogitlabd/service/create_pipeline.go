package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (s *Service) CreatePipeline(
	x context.Context,
	instance string,
	project int64,
	reference string,
	v []*gitlab.PipelineVariableOptions,
) (*pipeline_detail.Detail, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*pipeline_detail.Detail, error) {
			return c.CreatePipeline(y, project, reference, v)
		},
	)
}
