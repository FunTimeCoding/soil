package service

import "github.com/funtimecoding/soil/pkg/source/index"

func (s *Service) workspace(root string) *index.Workspace {
	return s.workspaces.Workspace(root)
}
