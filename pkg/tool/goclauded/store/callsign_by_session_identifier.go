package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) CallsignBySessionIdentifier(d string) (string, error) {
	var i session.Session
	result := s.database.Where(
		"identifier = ?",
		d,
	).Limit(1).Find(&i)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", nil
	}

	return i.CallsignValue(), nil
}
