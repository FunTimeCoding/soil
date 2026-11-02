package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) CommitDiff(
	x context.Context,
	instance string,
	project int64,
	sha string,
	limit int64,
) ([]*diff.Diff, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*diff.Diff, error) {
			return c.CommitDiff(y, project, sha, limit)
		},
	)
}
