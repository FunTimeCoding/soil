package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) StampPrompt(
	sessionIdentifier string,
	at time.Time,
) error {
	return s.database.Model(session.Stub()).Where(
		"identifier = ?",
		sessionIdentifier,
	).Update("last_prompt_at", at).Error
}
