package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) SweepCompleteTimeout(cutoff time.Time) []session.Session {
	var sessions []session.Session
	errors.PanicOnError(
		s.database.
			Where(
				"callsign IS NOT NULL AND callsign != '' AND last_seen < ? AND (last_active_at IS NULL OR last_active_at < ?) AND topic = '' AND timed_out = ''",
				cutoff,
				cutoff,
			).
			Find(&sessions).Error,
	)

	for _, e := range sessions {
		errors.PanicOnError(
			s.database.Model(session.Stub()).
				Where("callsign = ?", e.CallsignValue()).
				Update("timed_out", constant.CompleteTimeout).Error,
		)
	}

	return sessions
}
