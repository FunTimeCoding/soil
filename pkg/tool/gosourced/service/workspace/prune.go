package workspace

import "github.com/funtimecoding/soil/pkg/source/index"

func (s *Service) Prune() []string {
	return index.Prune(s.indexDirectory)
}
