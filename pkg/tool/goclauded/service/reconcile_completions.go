package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"time"
)

func (s *Service) ReconcileCompletions() {
	s.backfillCompletionSequences()
	existing := s.completionIndexer.Existing()
	pushed := 0
	skipped := 0

	for _, c := range s.store.ListCompletions() {
		slug, metadata, e := s.sessionMetadata(c.SessionIdentifier)
		errors.PanicOnError(e)

		if slug == "" {
			continue
		}

		name := fmt.Sprintf("%s/%d", slug, c.Sequence)

		if existing[name] == store.HashContent(c.Summary) {
			skipped++

			continue
		}

		metadata["created_at"] = c.CreatedAt.Format(time.RFC3339)
		s.completionIndexer.MustPush(name, c.Summary, metadata)
		pushed++
	}

	s.logger.Structured(
		"reconcile_completions",
		"pushed",
		pushed,
		"skipped",
		skipped,
	)
}
