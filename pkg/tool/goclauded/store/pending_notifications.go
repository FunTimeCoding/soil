package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"

func (s *Store) PendingNotifications(callsign string) ([]notification.Notification, error) {
	var result []notification.Notification

	if e := s.database.Where(
		"callsign = ? AND consumed = ?",
		callsign,
		false,
	).Order("created_at").Find(&result).Error; e != nil {
		return nil, e
	}

	if len(result) > 0 {
		if e := s.database.Model(notification.Stub()).Where(
			"callsign = ? AND consumed = ?",
			callsign,
			false,
		).Update("consumed", true).Error; e != nil {
			return nil, e
		}
	}

	return result, nil
}
