package service

import "github.com/funtimecoding/soil/pkg/gitlab/face"

func (s *Service) Trace(
	instance string,
	project int64,
	jobIdentifier int64,
) (string, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (string, error) {
			return c.Trace(project, jobIdentifier)
		},
	)
}
