package store

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/subscription"
)

func (s *Store) SubscriptionByName(
	name string,
) (*subscription.Subscription, error) {
	var result subscription.Subscription
	e := s.database.Where("name = ?", name).First(&result).Error

	if relational.NotFound(e) {
		return nil, not_found.New(constant.Subscription, name)
	}

	if e != nil {
		return nil, e
	}

	return &result, nil
}
