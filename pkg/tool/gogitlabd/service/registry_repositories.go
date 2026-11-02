package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/registry_repository"
)

func (s *Service) RegistryRepositories(
	x context.Context,
	instance string,
	project int64,
	panicOnForbidden bool,
) ([]*registry_repository.Repository, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*registry_repository.Repository, error) {
			return c.RegistryRepositories(y, project, panicOnForbidden)
		},
	)
}
