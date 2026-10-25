package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Store) ReleaseHeldCruise(
	sessionIdentifier string,
	limit int,
) (int64, error) {
	var identifiers []uint

	if e := s.database.Model(queue.Stub()).Where(
		"session_identifier = ? AND delivery = ? AND immediate = ? AND consumed_at IS NULL",
		sessionIdentifier,
		constant.DeliveryCruise,
		false,
	).Order("created_at").Limit(limit).Pluck(
		constant.Identifier,
		&identifiers,
	).Error; e != nil {
		return 0, e
	}

	if len(identifiers) == 0 {
		return 0, nil
	}

	result := s.database.Model(queue.Stub()).Where(
		"identifier IN ?",
		identifiers,
	).Updates(
		map[string]any{constant.Immediate: true, "immediate_at": s.clock()},
	)

	return result.RowsAffected, result.Error
}
