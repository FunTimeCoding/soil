package store

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Store) ListRootMemories(
	memoryType string,
	tag string,
	scope string,
	activeOnly bool,
) ([]record.MemorySummary, error) {
	roots := true

	return s.queryMemories(memoryType, tag, scope, activeOnly, &roots)
}
