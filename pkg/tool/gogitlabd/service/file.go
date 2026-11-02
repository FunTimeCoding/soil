package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
)

func (s *Service) File(
	instance string,
	project int64,
	branch string,
	name string,
) (*file.File, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*file.File, error) {
			return c.File(project, branch, name)
		},
	)
}
