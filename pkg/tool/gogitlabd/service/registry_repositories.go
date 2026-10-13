package service

import "github.com/funtimecoding/soil/pkg/gitlab/registry_repository"

func (s *Service) RegistryRepositories(
	instance string,
	project int64,
	panicOnForbidden bool,
) ([]*registry_repository.Repository, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.RegistryRepositories(project, panicOnForbidden)
}
