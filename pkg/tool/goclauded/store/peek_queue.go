package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Store) PeekQueue(
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
	).Order(
		"created_at",
	).Find(
		&result,
	).Error; e != nil {
		return nil, e
	}

	return result, nil
}
