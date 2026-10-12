package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Service) DeleteRegistryPath(
	project int64,
	path string,
) ([]string, error) {
	list, e := s.client.RegistryRepositories(project, false)

	if e != nil {
		return nil, e
	}

	var deleted []string

	for _, target := range []string{path, join.Slash([]string{path, "cache"})} {
		for _, repository := range list {
			if repository.Path != target {
				continue
			}

			if f := s.client.DeleteRegistryRepository(
				project,
				repository.Identifier,
			); f != nil {
				return deleted, f
			}

			deleted = append(deleted, target)

			break
		}
	}

	if len(deleted) == 0 {
		return nil, not_found.New("registry repository", path)
	}

	return deleted, nil
}
