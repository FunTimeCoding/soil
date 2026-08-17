package service

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Service) ColdBackfillAllSessions() *BackfillResult {
	r := &BackfillResult{}
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
