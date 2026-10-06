package store

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Store) ListMemories(
	memoryType string,
	tag string,
	scope string,
	activeOnly bool,
) ([]record.MemorySummary, error) {
	return s.queryMemories(memoryType, tag, scope, activeOnly, nil)
}
