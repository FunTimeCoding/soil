package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Service) DeleteRegistryPath(
	instance string,
	project int64,
	path string,
) ([]string, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]string, error) {
			list, e := c.RegistryRepositories(project, false)

			if e != nil {
				return nil, e
			}

			var deleted []string

			for _, target := range []string{
				path,
				join.Slash([]string{path, "cache"}),
			} {
				for _, repository := range list {
					if repository.Path != target {
						continue
					}

					if f := c.DeleteRegistryRepository(
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
		},
	)
}
