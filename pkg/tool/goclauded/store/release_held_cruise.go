package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Store) ReleaseHeldCruise(sessionIdentifier string) (int64, error) {
	result := s.database.Model(queue.Stub()).Where(
		"session_identifier = ? AND kind IN ? AND immediate = ? AND consumed_at IS NULL",
		sessionIdentifier,
		constant.CruiseKinds,
		false,
	).Updates(
		map[string]any{
			constant.Immediate:     true,
			"immediate_at":         s.clock(),
			constant.ReleasedColumn: true,
		},
	)

	return result.RowsAffected, result.Error
}
