package service

import "github.com/funtimecoding/soil/pkg/gitlab/discussion"

func (s *Service) MergeRequestDiscussions(
	project int64,
	identifier int64,
) ([]*discussion.Discussion, error) {
	return s.client.MergeRequestDiscussions(project, identifier)
}
