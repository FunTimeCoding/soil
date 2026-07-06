package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (s *Service) ReconcileCompletions() {
	s.backfillCompletionSequences()

	for _, c := range s.store.ListCompletions() {
		slug, metadata, e := s.sessionMetadata(c.SessionIdentifier)
		errors.PanicOnError(e)

		if slug == "" {
			continue
		}

		metadata["created_at"] = c.CreatedAt.Format(time.RFC3339)
		s.completionIndexer.MustPush(
			fmt.Sprintf("%s/%d", slug, c.Sequence),
			c.Summary,
			metadata,
		)
	}
}
