package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func (s *Service) MustReindexMemory(m *record.Memory) {
	if !s.indexable(m) {
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
