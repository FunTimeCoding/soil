package store

import (
	"gorm.io/gorm"
	"time"
)

type Store struct {
	mapper *gorm.DB
	clock  func() time.Time
}
