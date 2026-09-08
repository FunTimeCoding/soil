package store

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"

func (s *Store) ByCallsign(
	callsign string,
) ([]subscription.Subscription, error) {
	var result []subscription.Subscription

	return result, s.database.
		Where("callsign = ?", callsign).
		Order("last_event_at DESC").
		Find(&result).Error
}
