package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) CheckConsistency() {
	databaseSessions, e := s.store.AllSessions(0, 0)
	errors.PanicOnError(e)
	databaseSet := make(map[string]bool, len(databaseSessions))

	for _, i := range databaseSessions {
		databaseSet[i.Identifier] = true
	}

	for _, identifier := range s.cache.Keys() {
		if databaseSet[identifier] {
			continue
		}

		s.store.CreateDiscovered(identifier)
		s.RefreshFromCache(identifier)
		s.logger.Structured(
			"consistency_discovered",
			constant.Identifier,
			identifier,
		)
	}
}
