package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) indexable(m *store.Memory) bool {
	return m.IsActive && !s.isHidden(m.Tags)
}
