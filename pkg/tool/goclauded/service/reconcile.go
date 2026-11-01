package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
)

func (s *Service) Reconcile(context.Context) []error {
	summaries := s.ReconcileSummaries()

	for _, e := range summaries {
		if unavailable.Is(e) {
			return []error{e}
		}
	}

	return append(summaries, s.ReconcileCompletions()...)
}
