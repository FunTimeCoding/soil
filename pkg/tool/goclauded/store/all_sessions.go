package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) AllSessions(
	limit int,
	offset int,
) ([]session.Session, error) {
	var result []session.Session
	q := s.database.Order("last_seen DESC")

	if limit > 0 {
		q = q.Limit(limit)
	}

	if offset > 0 {
		q = q.Offset(offset)
	}

	return result, q.Find(&result).Error
}
