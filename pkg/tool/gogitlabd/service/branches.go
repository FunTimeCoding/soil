package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) Branches(
	instance string,
	project int64,
) ([]*branch.Branch, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*branch.Branch, error) {
			return c.Branches(project)
		},
	)
}
