package service

import "github.com/funtimecoding/soil/pkg/gitlab/face"

func (s *Service) DeleteProjectVariable(
	instance string,
	project int64,
	key string,
) error {
	return s.attempt(
		instance,
		func(c face.GitLabClient) error {
			return c.DeleteProjectVariable(project, key)
		},
	)
}
