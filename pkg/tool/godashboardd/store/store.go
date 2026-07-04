package store

import "gorm.io/gorm"

type Store struct {
	mapper *gorm.DB
}
