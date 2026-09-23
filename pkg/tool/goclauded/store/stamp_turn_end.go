package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) StampTurnEnd(
	sessionIdentifier string,
	at time.Time,
) error {
	return s.database.Model(session.Stub()).Where(
		"identifier = ?",
		sessionIdentifier,
	).Update("last_turn_end_at", at).Error
}
