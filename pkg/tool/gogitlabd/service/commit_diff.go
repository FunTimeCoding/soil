package service

import "github.com/funtimecoding/soil/pkg/gitlab/diff"

func (s *Service) CommitDiff(
	project int64,
	sha string,
	limit int64,
) ([]*diff.Diff, error) {
	return s.client.CommitDiff(project, sha, limit)
}
