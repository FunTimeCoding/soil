package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func (s *Store) countImmediateQueue(
	sessionIdentifier string,
	callsign string,
	since time.Time,
) (int64, error) {
	condition, arguments := sessionKeyMatch(sessionIdentifier, callsign)
	var pushed int64

	if e := s.database.Model(queue.Stub()).Where(
		condition,
		arguments...,
	).Where(
		"immediate = ? AND released = ?",
		true,
		false,
	).Where(
		"COALESCE(immediate_at, created_at) > ?",
		since,
	).Count(&pushed).Error; e != nil {
		return 0, e
	}

	var released int64

	if e := s.database.Model(queue.Stub()).Where(
		condition,
		arguments...,
	).Where(
		"released = ? AND immediate_at > ?",
		true,
		since,
	).Distinct("immediate_at").Count(&released).Error; e != nil {
		return 0, e
	}

	return pushed + released, nil
}
