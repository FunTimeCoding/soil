package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Service) pushImmediate(e *queue.Entry) (bool, error) {
	recent, f := s.store.CountImmediateQueue(
		e.SessionIdentifier,
		e.Callsign,
		s.clock().Add(-constant.ImmediateWindow),
	)

	if f != nil {
		return false, f
	}

	e.Immediate = recent < constant.ImmediateLimit

	return e.Immediate, s.pushEntry(e)
}
