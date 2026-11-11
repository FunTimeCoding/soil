package document

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
	current := map[string]bool{}
	named := map[string]bool{}

	for _, m := range s.store.ListSummaries() {
		slug, metadata, f := s.SessionMetadata(m.SessionIdentifier)
		errors.PanicOnError(f)

		if slug == "" {
			continue
		}

		current[slug] = true
		named[m.SessionIdentifier] = true

		if existing[slug].Hash == store.HashContent(m.Body) {
			skipped++

			continue
		}

		if g := s.PushSummary(slug, m.Body, metadata); g != nil {
			failed = append(failed, g)

			if connection.Classify(g) != nil {
				s.metric.CountReconcile(
					constant.SummaryCollection,
					pushed,
					len(failed),
				)

				return failed
			}

			continue
		}

		pushed++
	}

	swept, g := s.sweep(constant.SummaryCollection, existing, current, named)
	failed = append(failed, g...)
	s.metric.CountReconcile(constant.SummaryCollection, pushed, len(failed))
	s.metric.CountSwept(constant.SummaryCollection, swept)

	if len(g) > 0 && connection.Classify(g[len(g)-1]) != nil {
		return failed
	}

	s.metric.SetIndexed(constant.SummaryCollection, pushed+skipped)

	return failed
}
