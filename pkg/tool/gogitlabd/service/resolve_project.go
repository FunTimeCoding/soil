package service

import "github.com/funtimecoding/soil/pkg/gitlab/face"

func (s *Service) ResolveProject(
	instance string,
	identifier string,
) (int64, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (int64, error) {
			return c.ResolveProject(identifier)
		},
	)
}
