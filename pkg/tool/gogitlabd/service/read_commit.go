package service

import "github.com/funtimecoding/soil/pkg/gitlab/commit"

func (s *Service) ReadCommit(
	instance string,
	project int64,
	sha string,
) (*commit.Commit, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.ReadCommit(project, sha)
}
