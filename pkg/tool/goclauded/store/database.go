package store

import "gorm.io/gorm"

func (s *Store) Database() *gorm.DB {
	return s.database
}
