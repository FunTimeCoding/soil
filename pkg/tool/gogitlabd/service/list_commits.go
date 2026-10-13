package service

import "github.com/funtimecoding/soil/pkg/gitlab/commit"

func (s *Service) ListCommits(
	instance string,
	project int64,
	reference string,
	limit int64,
) ([]*commit.Commit, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.ListCommits(project, reference, limit)
}
