package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) syncIndex(m *store.Memory) error {
	if !s.indexable(m) {
		return s.indexer.Delete(
			ScopeCollection(m.Scope),
			memoryPath(m.Identifier),
		)
	}

	return s.indexer.Push(
		ScopeCollection(m.Scope),
		memoryPath(m.Identifier),
		m.Content,
		pushMetadata(m),
	)
}
