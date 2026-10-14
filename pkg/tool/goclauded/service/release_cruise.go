package service

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (s *Service) ReleaseCruise() {
	held, e := s.store.HeldCruiseSessions()
	errors.PanicOnError(e)
	released := false

	for _, identifier := range held {
		wakes, f := s.cruiseWakes(identifier)
		errors.PanicOnError(f)

		if !wakes {
			continue
		}

		recent, g := s.store.CountImmediateQueue(
			identifier,
			"",
			s.clock().Add(-constant.ImmediateWindow),
		)
		errors.PanicOnError(g)
		budget := constant.ImmediateLimit - int(recent)

		if budget <= 0 {
			continue
		}

		count, h := s.store.ReleaseHeldCruise(identifier, budget)
		errors.PanicOnError(h)

		if count == 0 {
			continue
		}

		errors.PanicOnError(s.store.MarkPromoted(identifier))
		released = true
	}

	if released {
		s.notify()
	}
}
