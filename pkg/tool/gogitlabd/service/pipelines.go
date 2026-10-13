package service

import "github.com/funtimecoding/soil/pkg/gitlab/pipeline"

func (s *Service) Pipelines(
	instance string,
	project int64,
	reference string,
	status string,
	limit int64,
) ([]*pipeline.Pipeline, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Pipelines(project, reference, status, limit)
}
