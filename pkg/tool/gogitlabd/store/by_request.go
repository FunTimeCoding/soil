package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Store) ByRequest(
	k *watched_request.Key,
) ([]subscription.Subscription, error) {
	var result []subscription.Subscription

	return result, s.database.
		Where(
			"instance = ? AND project = ? AND request = ?",
			k.Instance,
			k.Project,
			k.Request,
		).
		Order("identifier").
		Find(&result).Error
}
