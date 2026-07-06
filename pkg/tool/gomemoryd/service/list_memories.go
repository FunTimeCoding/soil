package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) ListMemories(
	memoryType string,
	tag string,
	activeOnly bool,
) ([]store.MemorySummary, error) {
	return s.store.ListMemories(memoryType, tag, activeOnly)
}
