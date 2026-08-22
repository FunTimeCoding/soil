package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) FindSession(identifier string) (*session.Session, bool, error) {
	var i session.Session
	result := s.database.Where("identifier = ?", identifier).Limit(1).Find(&i)

	if result.Error != nil {
		return nil, false, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, false, nil
	}

	return &i, true, nil
}
