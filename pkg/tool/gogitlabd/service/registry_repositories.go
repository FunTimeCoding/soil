package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
)

func (s *Service) RegistryRepositories(
	instance string,
	project int64,
	panicOnForbidden bool,
) ([]*registry_repository.Repository, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*registry_repository.Repository, error) {
			return c.RegistryRepositories(project, panicOnForbidden)
		},
	)
}
