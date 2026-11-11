package document

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
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
	current := map[string]bool{}
	named := map[string]bool{}

	for _, c := range s.store.ListCompletions() {
		slug, metadata, f := s.SessionMetadata(c.SessionIdentifier)
		errors.PanicOnError(f)

		if slug == "" {
			continue
		}

		path := fmt.Sprintf("%s/%d", slug, c.Sequence)
		current[path] = true
		named[c.SessionIdentifier] = true

		if indexedCurrent(
			existing,
			path,
			store.HashContent(c.Summary),
			completionMetadata(metadata, &c),
		) {
			skipped++

			continue
		}

		if g := s.PushCompletion(slug, &c, c.Summary, metadata); g != nil {
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

	swept, g := s.sweep(constant.CompletionCollection, existing, current, named)
	failed = append(failed, g...)
	s.metric.CountReconcile(constant.CompletionCollection, pushed, len(failed))
	s.metric.CountSwept(constant.CompletionCollection, swept)

	if len(g) > 0 && connection.Classify(g[len(g)-1]) != nil {
		return failed
	}

	s.metric.SetIndexed(constant.CompletionCollection, pushed+skipped)

	return failed
}
