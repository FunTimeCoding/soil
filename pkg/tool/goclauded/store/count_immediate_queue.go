package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func (s *Store) CountImmediateQueue(
	sessionIdentifier string,
	callsign string,
	since time.Time,
) (int64, error) {
	condition, arguments := sessionKeyMatch(sessionIdentifier, callsign)
	var result int64

	return result, s.database.Model(queue.Stub()).Where(
		condition,
		arguments...,
	).Where(
		"immediate = ?",
		true,
	).Where(
		"created_at > ?",
		since,
	).Count(&result).Error
}
