package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) CreateBranch(
	instance string,
	project int64,
	name string,
	reference string,
) (*branch.Branch, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*branch.Branch, error) {
			return c.CreateBranch(project, name, reference)
		},
	)
}
