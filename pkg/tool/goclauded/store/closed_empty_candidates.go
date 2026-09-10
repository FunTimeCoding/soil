package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) ClosedEmptyCandidates() ([]session.Session, error) {
	var result []session.Session

	return result, s.database.Where(
		"closed_at IS NOT NULL AND turn_count = 0 AND lines = 0",
	).Find(&result).Error
}
