package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) StaleCallsignSessions(
	cutoff time.Time,
) ([]session.Session, error) {
	var result []session.Session

	return result, s.database.Where(
		"callsign IS NOT NULL AND callsign != '' AND last_seen < ?",
		cutoff,
	).Find(&result).Error
}
