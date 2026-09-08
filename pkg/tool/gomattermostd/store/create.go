package store

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"

func (s *Store) Create(v *subscription.Subscription) error {
	return s.database.Create(v).Error
}
