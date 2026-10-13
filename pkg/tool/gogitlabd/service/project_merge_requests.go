package service

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (s *Service) ProjectMergeRequests(
	instance string,
	project int64,
	state string,
) ([]*merge_request.Request, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.ProjectMergeRequests(project, state)
}
