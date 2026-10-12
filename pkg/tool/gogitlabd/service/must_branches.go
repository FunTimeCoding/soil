package service

import "github.com/funtimecoding/soil/pkg/gitlab/branch"

func (s *Service) MustBranches(project int64) []*branch.Branch {
	return s.client.MustBranches(project)
}
