package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) CompleteTask(name string) (string, error) {
	var i session.Session
	result := s.database.Where(
		"callsign = ?",
		name,
	).Limit(1).Find(&i)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", nil
	}

	return i.Topic, s.database.Model(session.Stub()).Where(
		"callsign = ?",
		name,
	).Updates(
		map[string]any{
			"topic":     "",
			"files":     "",
			"last_seen": s.clock(),
		},
	).Error
}
