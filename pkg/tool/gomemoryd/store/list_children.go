package store

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Store) ListChildren(parentIdentifier int64) ([]record.MemorySummary, error) {
	return s.listMemoriesWithParent(parentIdentifier)
}
