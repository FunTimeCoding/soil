package store

func (s *Store) ListMemories(
	memoryType string,
	tag string,
	scope string,
	activeOnly bool,
) ([]MemorySummary, error) {
	return s.queryMemories(memoryType, tag, scope, activeOnly, nil)
}
