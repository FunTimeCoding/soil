package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) StaleEmptyCandidates(
	cutoff time.Time,
) ([]session.Session, error) {
	var result []session.Session

	return result, s.database.Where(
		"closed_at IS NULL AND turn_count = 0 AND lines = 0 AND last_seen < ? AND (last_active_at IS NULL OR last_active_at < ?)",
		cutoff,
		cutoff,
	).Find(&result).Error
}
