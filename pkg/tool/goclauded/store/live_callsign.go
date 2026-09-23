package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"time"
)

func (s *Store) LiveCallsign(
	sessionIdentifier string,
	since time.Time,
) (string, error) {
	var i session.Session
	result := s.database.Where(
		"identifier = ?",
		sessionIdentifier,
	).Limit(1).Find(&i)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 || i.LastPromptAt == nil {
		return "", nil
	}

	if !i.LastPromptAt.After(since) {
		return "", nil
	}

	return i.CallsignValue(), nil
}
