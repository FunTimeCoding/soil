package service

import "github.com/funtimecoding/soil/pkg/gitlab/diff"

func (s *Service) MergeRequestDiffs(
	project int64,
	identifier int64,
) ([]*diff.Diff, error) {
	return s.client.MergeRequestDiffs(project, identifier)
}
