package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"

func (s *Store) CompletionsBySession(
	sessionIdentifier string,
) ([]completion.Completion, error) {
	var result []completion.Completion

	return result, s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Order("created_at ASC").Find(&result).Error
}
