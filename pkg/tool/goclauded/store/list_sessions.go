package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) ListSessions() ([]session.Session, error) {
	var result []session.Session

	return result, s.database.Where(
		"last_seen > ? AND callsign IS NOT NULL",
		s.clock().Add(-1*time.Hour),
	).Order(constant.Callsign).Find(&result).Error
}
