package workspace

import "github.com/funtimecoding/soil/pkg/source/index"

func (s *Service) Workspace(root string) *index.Workspace {
	return s.workspaces.Workspace(root)
}
