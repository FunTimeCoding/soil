package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
)

func (s *Service) DeliverImpressions() {
	pending, e := s.store.PendingImpressions()
	errors.PanicOnError(e)

	for _, p := range pending {
		f := s.saveImpression(
			context.Background(),
			p.Content,
			p.Source,
			p.CapturedAt,
		)

		if unavailable.Is(f) {
			return
		}

		if f != nil {
			s.reportRefusedImpression(f, p.Content, p.Source, p.CapturedAt)
		}

		errors.PanicOnError(s.store.DeleteImpression(p.Identifier))
	}
}
