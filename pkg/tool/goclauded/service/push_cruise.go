package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Service) pushCruise(e *queue.Entry) (bool, error) {
	wakes, f := s.cruiseWakes(e.SessionIdentifier)

	if f != nil {
		return false, f
	}

	if !wakes {
		return false, s.pushEntry(e)
	}

	delivered, g := s.pushImmediate(e)

	if g != nil || !delivered {
		return delivered, g
	}

	return true, s.store.MarkPromoted(e.SessionIdentifier)
}
