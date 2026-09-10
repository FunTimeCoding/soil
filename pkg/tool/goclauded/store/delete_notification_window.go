package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"
	"time"
)

func (s *Store) DeleteNotificationWindow(
	callsign string,
	from time.Time,
	to time.Time,
) (int64, error) {
	result := s.database.Where(
		"callsign = ? AND created_at >= ? AND created_at <= ?",
		callsign,
		from.UTC(),
		to.UTC(),
	).Delete(notification.Stub())

	return result.RowsAffected, result.Error
}
