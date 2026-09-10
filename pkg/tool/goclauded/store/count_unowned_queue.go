package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) CountUnownedQueue() (int64, error) {
	var result int64

	return result, s.database.Model(queue.Stub()).Where(
		"consumed_at IS NULL AND callsign NOT IN (?)",
		s.database.Model(session.Stub()).Select(constant.Callsign).Where(
			"callsign IS NOT NULL AND callsign != ''",
		),
	).Count(&result).Error
}
