package service

import "github.com/funtimecoding/soil/pkg/gitlab/diff"

func (s *Service) MergeRequestDiffs(
	instance string,
	project int64,
	identifier int64,
) ([]*diff.Diff, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.MergeRequestDiffs(project, identifier)
}
