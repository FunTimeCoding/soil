package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) SetListening(
	name string,
	listening bool,
) error {
	return s.database.Model(session.Stub()).Where(
		"callsign = ?",
		name,
	).Updates(
		map[string]any{
			"listening": listening,
			"last_seen": s.clock(),
		},
	).Error
}
