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

	cacheKeys := s.cache.Keys()
	cacheSet := make(map[string]bool, len(cacheKeys))

	for _, identifier := range cacheKeys {
		cacheSet[identifier] = true
	}

	for identifier := range cacheSet {
		if !databaseSet[identifier] {
			s.store.CreateDiscovered(identifier)
			s.RefreshFromCache(identifier)
			s.logger.Structured(
				"consistency_discovered",
				constant.Identifier,
				identifier,
			)
		}
	}

	for _, i := range databaseSessions {
		if !cacheSet[i.Identifier] {
			s.logger.Structured(
				"consistency_missing_jsonl",
				constant.Identifier,
				i.Identifier,
				constant.SessionName,
				i.Name,
				constant.Alias,
				i.Alias,
			)
		}
	}
}
