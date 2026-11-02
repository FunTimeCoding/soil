package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) MergeRequestDiscussions(
	instance string,
	project int64,
	identifier int64,
) ([]*discussion.Discussion, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*discussion.Discussion, error) {
			return c.MergeRequestDiscussions(project, identifier)
		},
	)
}
