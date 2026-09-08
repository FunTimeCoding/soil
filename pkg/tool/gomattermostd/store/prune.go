package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"time"
)

func (s *Store) Prune(cutoff time.Time) ([]subscription.Subscription, error) {
	var result []subscription.Subscription
	e := s.database.
		Where("last_event_at < ?", cutoff).
		Find(&result).Error

	if e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return result, nil
	}

	return result, s.database.
		Where("last_event_at < ?", cutoff).
		Delete(subscription.Stub()).Error
}
