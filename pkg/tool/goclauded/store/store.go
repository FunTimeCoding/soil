package store

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type Store struct {
	database *gorm.DB
	clock    func() time.Time
	mutex    sync.Mutex
}
