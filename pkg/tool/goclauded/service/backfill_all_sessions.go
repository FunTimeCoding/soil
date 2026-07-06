package service

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Service) BackfillAllSessions() *BackfillResult {
	r := &BackfillResult{}
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
