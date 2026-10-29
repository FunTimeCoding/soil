package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func (s *Service) ReconcileSummaries() []error {
	existing, e := s.indexed(constant.SummaryCollection)

	if e != nil {
		return []error{e}
	}

	pushed := 0
	skipped := 0
	var failed []error

	for _, m := range s.store.ListSummaries() {
		slug, metadata, f := s.sessionMetadata(m.SessionIdentifier)
		errors.PanicOnError(f)

		if slug == "" {
			continue
		}

		if existing[slug] == store.HashContent(m.Body) {
			skipped++

			continue
		}

		if g := s.pushSummary(slug, m.Body, metadata); g != nil {
			failed = append(failed, g)

			if connection.Classify(g) != nil {
				return failed
			}

			continue
		}

		pushed++
	}

	if pushed > 0 {
		s.logger.Structured(
			"reconcile_summaries",
			"pushed",
			pushed,
			"skipped",
			skipped,
		)
	}

	return failed
}
