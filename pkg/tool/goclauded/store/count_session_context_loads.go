package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"

func (s *Store) CountSessionContextLoads(
	sessionIdentifier string,
) (int64, error) {
	var result int64

	return result, s.database.Model(context_load.Stub()).Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Count(&result).Error
}
