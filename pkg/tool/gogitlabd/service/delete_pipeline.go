package service

import "github.com/funtimecoding/soil/pkg/gitlab/face"

func (s *Service) DeletePipeline(
	instance string,
	project int64,
	identifier int64,
) error {
	return s.attempt(
		instance,
		func(c face.GitLabClient) error {
			return c.DeletePipeline(project, identifier)
		},
	)
}
