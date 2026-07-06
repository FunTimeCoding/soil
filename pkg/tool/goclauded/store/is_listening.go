package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) isListening(name string) (bool, error) {
	var i session.Session
	result := s.database.Where("callsign = ?", name).Limit(1).Find(&i)

	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	}

	return i.Listening, nil
}
