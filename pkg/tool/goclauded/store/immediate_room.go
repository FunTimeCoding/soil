package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func (s *Store) ImmediateRoom(
	sessionIdentifier string,
	callsign string,
	now time.Time,
) (bool, error) {
	recent, e := s.countImmediateQueue(
		sessionIdentifier,
		callsign,
		now.Add(-constant.ImmediateWindow),
	)

	if e != nil {
		return false, e
	}

	return recent < constant.ImmediateLimit, nil
}
