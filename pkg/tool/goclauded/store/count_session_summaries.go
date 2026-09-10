package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/summary"

func (s *Store) CountSessionSummaries(sessionIdentifier string) (int64, error) {
	var result int64

	return result, s.database.Model(summary.Stub()).Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Count(&result).Error
}
