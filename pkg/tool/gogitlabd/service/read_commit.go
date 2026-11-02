package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) ReadCommit(
	instance string,
	project int64,
	sha string,
) (*commit.Commit, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*commit.Commit, error) {
			return c.ReadCommit(project, sha)
		},
	)
}
