package message

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) releaseHeld(identifier string) (bool, error) {
	wakes, e := s.cruiseWakes(identifier)

	if e != nil || !wakes {
		return false, e
	}

	recent, f := s.store.CountImmediateQueue(
		identifier,
		"",
		s.clock().Add(-constant.ImmediateWindow),
	)

	if f != nil {
		return false, f
	}

	budget := constant.ImmediateLimit - int(recent)

	if budget <= 0 {
		return false, nil
	}

	count, g := s.store.ReleaseHeldCruise(identifier, budget)

	if g != nil || count == 0 {
		return false, g
	}

	return true, s.store.MarkPromoted(identifier)
}
