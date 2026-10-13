package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Service) DeleteRegistryPath(
	instance string,
	project int64,
	path string,
) ([]string, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	list, f := c.RegistryRepositories(project, false)

	if f != nil {
		return nil, f
	}

	var deleted []string

	for _, target := range []string{path, join.Slash([]string{path, "cache"})} {
		for _, repository := range list {
			if repository.Path != target {
				continue
			}

			if g := c.DeleteRegistryRepository(
				project,
				repository.Identifier,
			); g != nil {
				return deleted, g
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
