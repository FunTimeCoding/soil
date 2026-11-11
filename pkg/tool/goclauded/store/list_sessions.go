package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) ListSessions() ([]session.Session, error) {
	var result []session.Session

	return result, s.database.Where(
		"last_seen > ? AND callsign IS NOT NULL",
		s.clock().Add(-constant.InactivityTimeoutWindow),
	).Order(constant.Callsign).Find(&result).Error
}
