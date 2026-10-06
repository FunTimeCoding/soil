package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/backfill_result"
)

func (s *Service) BackfillAllSessions() *backfill_result.Result {
	r := backfill_result.New()
	sessions, e := s.store.AllSessions(0, 0)
	errors.PanicOnError(e)

	for _, e := range sessions {
		resolved := s.client.Resolve(e.Identifier)

		if resolved.Identifier == "" {
			r.Skipped++

			continue
		}

		s.EnrichSession(e.Identifier)
		r.Enriched++
	}

	return r
}
