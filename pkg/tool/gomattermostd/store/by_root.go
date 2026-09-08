package store

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"

func (s *Store) ByRoot(root string) ([]subscription.Subscription, error) {
	var result []subscription.Subscription

	return result, s.database.
		Where("root_identifier = ?", root).
		Find(&result).Error
}
