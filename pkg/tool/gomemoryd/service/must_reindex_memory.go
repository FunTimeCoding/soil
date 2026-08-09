package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func (s *Service) MustReindexMemory(m *store.Memory) {
	if s.isHidden(m.Tags) {
		errors.LogOnError(
			s.indexer.Delete(
				ScopeCollection(m.Scope),
				memoryPath(m.Identifier),
			),
		)

		return
	}

	s.indexer.MustPush(
		ScopeCollection(m.Scope),
		memoryPath(m.Identifier),
		m.Content,
		pushMetadata(m),
	)
}
