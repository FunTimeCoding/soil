package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"

func (s *Store) MarkPromoted(sessionIdentifier string) error {
	return s.database.Model(cruise_setting.Stub()).Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Update("promoted_at", s.clock().UTC()).Error
}
