package store

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"

func (s *Store) All() ([]subscription.Subscription, error) {
	var result []subscription.Subscription

	return result, s.database.Find(&result).Error
}
