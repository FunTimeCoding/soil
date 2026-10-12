package service

import "github.com/funtimecoding/soil/pkg/gitlab/branch"

func (s *Service) CreateBranch(
	project int64,
	name string,
	reference string,
) (*branch.Branch, error) {
	return s.client.CreateBranch(project, name, reference)
}
