package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Service) sweepEmptySessions() {
	closed, closedError := s.store.ClosedEmptyCandidates()
	errors.PanicOnError(closedError)
	stale, staleError := s.store.StaleEmptyCandidates(
		s.clock().Add(-24 * time.Hour),
	)
	errors.PanicOnError(staleError)
	swept := 0

	for _, i := range append(closed, stale...) {
		refused, refusalError := s.EmptyRefusal(&i)
		errors.PanicOnError(refusalError)

		if refused != nil {
			continue
		}

		if _, e := s.DeleteSession(i.Identifier, ""); e != nil {
			s.logger.Structured(
				"sweep_empty_refused",
				constant.Identifier,
				i.Identifier,
				"reason",
				e.Error(),
			)

			continue
		}

		swept++
		s.logger.Structured(
			"swept_empty_session",
			constant.Identifier,
			i.Identifier,
			constant.SessionName,
			i.Name,
			"reason",
			i.ClosedReason,
		)
	}

	if swept > 0 {
		s.notify()
	}
}
