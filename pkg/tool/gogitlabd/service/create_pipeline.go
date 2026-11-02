package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (s *Service) CreatePipeline(
	instance string,
	project int64,
	reference string,
	v []*gitlab.PipelineVariableOptions,
) (*pipeline_detail.Detail, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*pipeline_detail.Detail, error) {
			return c.CreatePipeline(project, reference, v)
		},
	)
}
