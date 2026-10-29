package store

import (
	"gorm.io/gorm"
	"time"
)

func New(
	d *gorm.DB,
	clock func() time.Time,
) *Store {
	return &Store{mapper: d, clock: clock}
}
