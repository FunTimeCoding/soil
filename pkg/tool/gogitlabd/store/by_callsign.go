package store

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"

func (s *Store) ByCallsign(
	callsign string,
) ([]subscription.Subscription, error) {
	var result []subscription.Subscription

	return result, s.database.
		Where("callsign = ?", callsign).
		Order("identifier").
		Find(&result).Error
}
