package store

func (s *Store) ListChildren(parentIdentifier int64) ([]MemorySummary, error) {
	return s.listMemoriesWithParent(parentIdentifier)
}
