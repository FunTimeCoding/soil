package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) MergeRequestDiscussions(
	x context.Context,
	instance string,
	project int64,
	identifier int64,
) ([]*discussion.Discussion, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*discussion.Discussion, error) {
			return c.MergeRequestDiscussions(y, project, identifier)
		},
	)
}
