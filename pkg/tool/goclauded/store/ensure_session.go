package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/ensure_result"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) EnsureSession(identifier string) (*ensure_result.Result, error) {
	var i session.Session
	result := s.database.Where(
		"identifier = ?",
		identifier,
	).Limit(1).Find(&i)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected > 0 {
		previousLastSeen := i.LastSeen

		if e := s.database.Model(&i).Updates(
			map[string]any{
				"last_seen":  s.clock(),
				"turn_count": i.TurnCount + 1,
				"timed_out":  "",
			},
		).Error; e != nil {
			return nil, e
		}

		return ensure_result.NewExisting(
			i.CallsignValue(),
			previousLastSeen,
		), nil
	}

	name, e := s.NextName()

	if e != nil {
		return nil, e
	}

	if name == "" {
		return ensure_result.Stub(), nil
	}

	if f := s.database.Create(
		session.NewRegistered(identifier, name),
	).Error; f != nil {
		return nil, f
	}

	return ensure_result.New(name), nil
}
