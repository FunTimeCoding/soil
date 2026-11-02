package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) MergeRequestDiffs(
	x context.Context,
	instance string,
	project int64,
	identifier int64,
) ([]*diff.Diff, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*diff.Diff, error) {
			return c.MergeRequestDiffs(y, project, identifier)
		},
	)
}
