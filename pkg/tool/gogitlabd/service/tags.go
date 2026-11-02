package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func (s *Service) Tags(
	instance string,
	project int64,
) ([]*tag.Tag, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*tag.Tag, error) {
			return c.Tags(project)
		},
	)
}
