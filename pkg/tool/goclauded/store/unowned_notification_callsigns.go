package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/notification"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) UnownedNotificationCallsigns() ([]string, error) {
	var result []string

	return result, s.database.Model(notification.Stub()).Where(
		"consumed = ? AND callsign NOT IN (?)",
		false,
		s.database.Model(session.Stub()).Select(constant.Callsign).Where(
			"callsign IS NOT NULL AND callsign != ''",
		),
	).Distinct().Order(constant.Callsign).Pluck(
		constant.Callsign,
		&result,
	).Error
}
