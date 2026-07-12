package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) UnnamedSessions(
	limit int,
	offset int,
) ([]session.Session, error) {
	var result []session.Session
	q := s.database.
		Where("alias IS NULL OR alias = ''").
		Order("session_timestamp")

	if limit > 0 {
		q = q.Limit(limit)
	}

	if offset > 0 {
		q = q.Offset(offset)
	}

	return result, q.Find(&result).Error
}
