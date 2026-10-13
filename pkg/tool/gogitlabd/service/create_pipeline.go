package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (s *Service) CreatePipeline(
	instance string,
	project int64,
	reference string,
	v []*gitlab.PipelineVariableOptions,
) (*pipeline_detail.Detail, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.CreatePipeline(project, reference, v)
}
