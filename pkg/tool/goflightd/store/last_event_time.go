package store

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"gorm.io/gorm"
	"time"
)

func (s *Store) LastEventTime() (*time.Time, error) {
	var result event.Event
	e := s.database.Order("time DESC").First(&result).Error

	if errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if e != nil {
		return nil, e
	}

	return &result.Time, nil
}
