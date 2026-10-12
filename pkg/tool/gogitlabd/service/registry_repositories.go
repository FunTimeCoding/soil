package service

import "github.com/funtimecoding/soil/pkg/gitlab/registry_repository"

func (s *Service) RegistryRepositories(
	project int64,
	panicOnForbidden bool,
) ([]*registry_repository.Repository, error) {
	return s.client.RegistryRepositories(project, panicOnForbidden)
}
