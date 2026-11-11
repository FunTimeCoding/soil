package document

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/stored_document"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func (s *Service) reconcileCollection(
	collection string,
	sourceType string,
	documents []*stored_document.Document,
) []error {
	existing, e := s.indexed(collection)

	if e != nil {
		return []error{e}
	}

	pushed := 0
	skipped := 0
	var failed []error
	current := map[string]bool{}
	named := map[string]bool{}

	for _, d := range documents {
		current[d.Path] = true
		named[d.SessionIdentifier] = true

		if indexedCurrent(
			existing,
			d.Path,
			store.HashContent(d.Body),
			d.Metadata,
		) {
			skipped++

			continue
		}

		if f := s.pushDocument(
			collection,
			sourceType,
			d.Path,
			d.Body,
			d.Metadata,
		); f != nil {
			failed = append(failed, f)

			if connection.Classify(f) != nil {
				s.metric.CountReconcile(collection, pushed, len(failed))

				return failed
			}

			continue
		}

		pushed++
	}

	swept, g := s.sweep(collection, existing, current, named)
	failed = append(failed, g...)
	s.metric.CountReconcile(collection, pushed, len(failed))
	s.metric.CountSwept(collection, swept)

	if len(g) > 0 && connection.Classify(g[len(g)-1]) != nil {
		return failed
	}

	s.metric.SetIndexed(collection, pushed+skipped)

	return failed
}
