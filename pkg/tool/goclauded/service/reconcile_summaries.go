package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func (s *Service) ReconcileSummaries() {
	existing := s.summaryIndexer.Existing()
	pushed := 0
	skipped := 0

	for _, m := range s.store.ListSummaries() {
		slug, metadata, e := s.sessionMetadata(m.SessionIdentifier)
		errors.PanicOnError(e)

		if slug == "" {
			continue
		}

		if existing[slug] == store.HashContent(m.Body) {
			skipped++

			continue
		}

		s.summaryIndexer.MustPush(slug, m.Body, metadata)
		pushed++
	}

	s.logger.Structured(
		"reconcile_summaries",
		"pushed",
		pushed,
		"skipped",
		skipped,
	)
}
