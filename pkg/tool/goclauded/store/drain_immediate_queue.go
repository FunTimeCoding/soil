package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Store) DrainImmediateQueue(
	sessionIdentifier string,
	callsign string,
) ([]queue.Entry, error) {
	condition, arguments := sessionKeyMatch(sessionIdentifier, callsign)
	var result []queue.Entry

	if e := s.database.Where(
		condition,
		arguments...,
	).Where(
		"consumed_at IS NULL",
	).Where(
		"immediate = ?",
		true,
	).Order(
		"created_at",
	).Find(
		&result,
	).Error; e != nil {
		return nil, e
	}

	if len(result) > 0 {
		if e := s.database.Model(queue.Stub()).Where(
			condition,
			arguments...,
		).Where(
			"consumed_at IS NULL",
		).Where(
			"immediate = ?",
			true,
		).Updates(
			map[string]any{"consumed": true, "consumed_at": s.clock()},
		).Error; e != nil {
			return nil, e
		}
	}

	return result, nil
}
