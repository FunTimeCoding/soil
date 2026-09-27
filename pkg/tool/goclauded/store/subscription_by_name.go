package store

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/subscription"
)

func (s *Store) SubscriptionByName(
	name string,
) (*subscription.Subscription, error) {
	var result subscription.Subscription

	if e := s.database.Where("name = ?", name).First(&result).Error; e != nil {
		return nil, not_found.New(constant.Subscription, name)
	}

	return &result, nil
}
