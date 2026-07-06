package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) GetSession(identifier string) (*session.Session, error) {
	var i session.Session
	result := s.database.Where("identifier = ?", identifier).Limit(1).Find(&i)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &i, nil
}
