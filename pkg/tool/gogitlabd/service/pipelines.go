package service

import "github.com/funtimecoding/soil/pkg/gitlab/pipeline"

func (s *Service) Pipelines(
	project int64,
	reference string,
	status string,
	limit int64,
) ([]*pipeline.Pipeline, error) {
	return s.client.Pipelines(project, reference, status, limit)
}
