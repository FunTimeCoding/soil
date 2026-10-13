package service

import "github.com/funtimecoding/soil/pkg/gitlab/diff"

func (s *Service) CommitDiff(
	instance string,
	project int64,
	sha string,
	limit int64,
) ([]*diff.Diff, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.CommitDiff(project, sha, limit)
}
