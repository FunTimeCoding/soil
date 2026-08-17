package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"

func (s *Store) ContextLoadsBySession(
	sessionIdentifier string,
) ([]context_load.Load, error) {
	var result []context_load.Load

	return result, s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Order("occurred_at, identifier").Find(&result).Error
}
