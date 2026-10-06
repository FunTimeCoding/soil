package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) ListMemories(
	memoryType string,
	tag string,
	scope string,
	activeOnly bool,
) ([]record.MemorySummary, error) {
	return s.store.ListMemories(memoryType, tag, scope, activeOnly)
}
