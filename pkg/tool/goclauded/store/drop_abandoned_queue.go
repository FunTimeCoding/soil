package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func (s *Store) DropAbandonedQueue(before time.Time) int64 {
	result := s.database.Where(
		`consumed_at IS NULL
		AND COALESCE(session_identifier, '') != ''
		AND (
			session_identifier NOT IN (SELECT identifier FROM session)
			OR session_identifier IN (
				SELECT identifier FROM session
				WHERE (closed_at IS NOT NULL AND closed_at < ?)
				OR (COALESCE(callsign, '') = '' AND last_seen < ?)
			)
		)`,
		before,
		before,
	).Delete(
		queue.Stub(),
	)

	return result.RowsAffected
}
