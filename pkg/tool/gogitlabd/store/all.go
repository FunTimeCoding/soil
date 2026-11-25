package store

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"

func (s *Store) All() ([]subscription.Subscription, error) {
	var result []subscription.Subscription

	return result, s.database.Order("identifier").Find(&result).Error
}
