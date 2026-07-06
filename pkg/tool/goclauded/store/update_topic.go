package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) UpdateTopic(
	name string,
	topic string,
	files string,
) error {
	return s.database.Model(session.Stub()).Where(
		"callsign = ?",
		name,
	).Updates(
		map[string]any{
			"topic":     topic,
			"files":     files,
			"last_seen": s.clock(),
		},
	).Error
}
