package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/diff"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) MergeRequestDiffs(
	instance string,
	project int64,
	identifier int64,
) ([]*diff.Diff, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*diff.Diff, error) {
			return c.MergeRequestDiffs(project, identifier)
		},
	)
}
