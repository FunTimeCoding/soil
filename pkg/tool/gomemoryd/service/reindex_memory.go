package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) ReindexMemory(m *record.Memory) error {
	return s.syncIndex(m)
}
