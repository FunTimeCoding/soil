package service

import "github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"

func (s *Service) Pipeline(
	instance string,
	project int64,
	identifier int64,
) (*pipeline_detail.Detail, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Pipeline(project, identifier)
}
