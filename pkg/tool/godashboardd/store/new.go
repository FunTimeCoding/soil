package store

import "gorm.io/gorm"

func New(m *gorm.DB) *Store {
	return &Store{mapper: m}
}
