package store

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) Announce(
	name string,
	topic string,
	files string,
) error {
	result := s.database.Model(session.Stub()).Where("callsign = ?", name).Updates(
		map[string]any{
			"topic":     topic,
			"files":     files,
			"last_seen": s.clock(),
		},
	)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return not_found.New(constant.Callsign, name)
	}

	return nil
}
