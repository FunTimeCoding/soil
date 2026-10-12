package service

import "github.com/funtimecoding/soil/pkg/gitlab/commit"

func (s *Service) ReadCommit(
	project int64,
	sha string,
) (*commit.Commit, error) {
	return s.client.ReadCommit(project, sha)
}
