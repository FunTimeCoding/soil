package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"

func (s *Store) CountSessionCompletions(
	sessionIdentifier string,
) (int64, error) {
	var result int64

	return result, s.database.Model(completion.Stub()).Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Count(&result).Error
}
