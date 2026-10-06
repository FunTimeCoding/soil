package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) indexable(m *record.Memory) bool {
	return m.IsActive && !s.isHidden(m.Tags)
}
