package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) ResolveByCallsign(c string) (string, error) {
	var i session.Session
	result := s.database.Where("callsign = ?", c).Limit(1).Find(&i)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", nil
	}

	return i.Identifier, nil
}
