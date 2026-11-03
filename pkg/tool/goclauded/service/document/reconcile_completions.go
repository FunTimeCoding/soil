package document

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"time"
)

func (s *Service) ReconcileCompletions() []error {
	s.backfillCompletionSequences()
	existing, e := s.indexed(constant.CompletionCollection)

	if e != nil {
		return []error{e}
	}

	pushed := 0
	skipped := 0
	var failed []error

	for _, c := range s.store.ListCompletions() {
		slug, metadata, f := s.SessionMetadata(c.SessionIdentifier)
		errors.PanicOnError(f)

		if slug == "" {
			continue
		}

		name := fmt.Sprintf("%s/%d", slug, c.Sequence)

		if existing[name] == store.HashContent(c.Summary) {
			skipped++

			continue
		}

		metadata["created_at"] = []string{c.CreatedAt.Format(time.RFC3339)}

		if g := s.pushDocument(
			constant.CompletionCollection,
			constant.CompletionSourceType,
			name,
			c.Summary,
			metadata,
		); g != nil {
			failed = append(failed, g)

			if connection.Classify(g) != nil {
				s.metric.CountReconcile(
					constant.CompletionCollection,
					pushed,
					len(failed),
				)

				return failed
			}

			continue
		}

		pushed++
	}

	s.metric.CountReconcile(constant.CompletionCollection, pushed, len(failed))
	s.metric.SetIndexed(constant.CompletionCollection, pushed+skipped)

	return failed
}
