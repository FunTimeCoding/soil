package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) ReadCommit(
	x context.Context,
	instance string,
	project int64,
	sha string,
) (*commit.Commit, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*commit.Commit, error) {
			return c.ReadCommit(y, project, sha)
		},
	)
}
