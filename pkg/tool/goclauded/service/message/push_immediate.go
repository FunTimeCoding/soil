package message

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Service) pushImmediate(e *queue.Entry) (bool, error) {
	room, f := s.store.ImmediateRoom(e.SessionIdentifier, e.Callsign, s.clock())

	if f != nil {
		return false, f
	}

	e.Immediate = room

	if e.Immediate {
		e.ImmediateAt = new(s.clock())
	}

	return e.Immediate, s.pushEntry(e)
}
