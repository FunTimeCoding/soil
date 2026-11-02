package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) CommitDiff(
	instance string,
	project int64,
	sha string,
	limit int64,
) ([]*diff.Diff, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*diff.Diff, error) {
			return c.CommitDiff(project, sha, limit)
		},
	)
}
