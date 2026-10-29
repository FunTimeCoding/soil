package store

import "gorm.io/gorm"

func New(
	m *gorm.DB,
	tableName string,
) *Store {
	return &Store{mapper: m, tableName: tableName}
}
