package service

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Service) ReconcileSummaries() {
	for _, m := range s.store.ListSummaries() {
		slug, metadata, e := s.sessionMetadata(m.SessionIdentifier)
		errors.PanicOnError(e)

		if slug == "" {
			continue
		}

		s.summaryIndexer.MustPush(slug, m.Body, metadata)
	}
}
