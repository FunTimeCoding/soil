package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (s *Service) CreatePipeline(
	project int64,
	reference string,
	v []*gitlab.PipelineVariableOptions,
) (*pipeline_detail.Detail, error) {
	return s.client.CreatePipeline(project, reference, v)
}
