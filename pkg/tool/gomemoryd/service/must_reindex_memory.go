package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Service) MustReindexMemory(m *store.Memory) {
	s.indexer.MustPush(
		ScopeCollection(m.Scope),
		memoryPath(m.Identifier),
		m.Content,
		pushMetadata(m),
	)
}
