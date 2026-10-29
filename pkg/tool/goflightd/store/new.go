package store

import "gorm.io/gorm"

func New(d *gorm.DB) *Store {
	return &Store{database: d}
}
