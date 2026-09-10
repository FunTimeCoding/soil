package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"

func (s *Store) PendingNotifications(
	sessionIdentifier string,
) ([]notification.Notification, error) {
	var result []notification.Notification

	if e := s.database.Where(
		"session_identifier = ? AND consumed = ?",
		sessionIdentifier,
		false,
	).Order(
		"created_at",
	).Find(
		&result,
	).Error; e != nil {
		return nil, e
	}

	if len(result) > 0 {
		if e := s.database.Model(notification.Stub()).Where(
			"session_identifier = ? AND consumed = ?",
			sessionIdentifier,
			false,
		).Update("consumed", true).Error; e != nil {
			return nil, e
		}
	}

	return result, nil
}
