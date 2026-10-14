package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Store) HeldCruiseSessions() ([]string, error) {
	var result []string

	return result, s.database.Model(queue.Stub()).Where(
		"delivery = ? AND immediate = ? AND consumed_at IS NULL",
		constant.DeliveryCruise,
		false,
	).Distinct().Pluck("session_identifier", &result).Error
}
