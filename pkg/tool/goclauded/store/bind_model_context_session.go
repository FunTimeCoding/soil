package store

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) BindModelContextSession(
	name string,
	modelContextSessionIdentifier string,
) error {
	result := s.database.Model(session.Stub()).Where("callsign = ?", name).Updates(
		map[string]any{"model_context_session": modelContextSessionIdentifier},
	)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return not_found.New(constant.Callsign, name)
	}

	return nil
}
