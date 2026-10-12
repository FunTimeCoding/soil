package service

import "github.com/funtimecoding/soil/pkg/gitlab/commit"

func (s *Service) ListCommits(
	project int64,
	reference string,
	limit int64,
) ([]*commit.Commit, error) {
	return s.client.ListCommits(project, reference, limit)
}
