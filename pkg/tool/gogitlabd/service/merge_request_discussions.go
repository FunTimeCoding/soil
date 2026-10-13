package service

import "github.com/funtimecoding/soil/pkg/gitlab/discussion"

func (s *Service) MergeRequestDiscussions(
	instance string,
	project int64,
	identifier int64,
) ([]*discussion.Discussion, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.MergeRequestDiscussions(project, identifier)
}
