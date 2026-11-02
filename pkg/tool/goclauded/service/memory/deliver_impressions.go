package memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
)

func (s *Service) DeliverImpressions() []error {
	pending, e := s.store.PendingImpressions()

	if e != nil {
		return []error{e}
	}

	for _, p := range pending {
		f := s.saveImpression(
			context.Background(),
			p.Content,
			p.Source,
			p.CapturedAt,
		)

		if unavailable.Is(f) {
			return nil
		}

		if f != nil {
			s.reportRefusedImpression(f, p.Content, p.Source, p.CapturedAt)
		}

		if g := s.store.DeleteImpression(p.Identifier); g != nil {
			return []error{g}
		}
	}

	return nil
}
