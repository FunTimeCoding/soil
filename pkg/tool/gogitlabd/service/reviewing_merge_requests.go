package service

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (s *Service) ReviewingMergeRequests(
	instance string,
	all bool,
) ([]*merge_request.Request, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.ReviewingMergeRequests(all)
}
