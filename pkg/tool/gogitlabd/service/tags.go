package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func (s *Service) Tags(
	x context.Context,
	instance string,
	project int64,
) ([]*tag.Tag, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*tag.Tag, error) {
			return c.Tags(y, project)
		},
	)
}
