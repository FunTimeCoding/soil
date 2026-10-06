package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/backfill_result"
)

func (s *Service) ColdBackfillAllSessions() *backfill_result.Result {
	r := backfill_result.New()
	sessions, e := s.store.AllSessions(0, 0)
	errors.PanicOnError(e)

	for _, entry := range sessions {
		resolved := s.client.Resolve(entry.Identifier)

		if resolved.Identifier == "" {
			r.Skipped++

			continue
		}

		s.cache.GetOrCreate(entry.Identifier).Reset()
		s.EnrichSession(entry.Identifier)
		r.Enriched++
	}

	return r
}
