package service

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
	existing, e := s.completionIndexer.Existing(constant.CompletionCollection)

	if e != nil {
		return []error{e}
	}

	pushed := 0
	skipped := 0
	var failed []error

	for _, c := range s.store.ListCompletions() {
		slug, metadata, f := s.sessionMetadata(c.SessionIdentifier)
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

		if g := s.completionIndexer.Push(
			constant.CompletionCollection,
			name,
			c.Summary,
			metadata,
		); g != nil {
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
			"reconcile_completions",
			"pushed",
			pushed,
			"skipped",
			skipped,
		)
	}

	return failed
}
