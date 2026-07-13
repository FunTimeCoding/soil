package store

import "gorm.io/gorm"

type Store struct {
	database *gorm.DB
}
