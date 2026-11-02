package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) ListCommits(
	instance string,
	project int64,
	reference string,
	limit int64,
) ([]*commit.Commit, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*commit.Commit, error) {
			return c.ListCommits(project, reference, limit)
		},
	)
}
