package service

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"

func (s *Service) MergeRequest(
	instance string,
	project int64,
	identifier int64,
) (*merge_request_detail.Detail, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.MergeRequest(project, identifier)
}
