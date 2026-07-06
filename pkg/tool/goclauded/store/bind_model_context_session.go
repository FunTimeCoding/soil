package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) BindModelContextSession(
	name string,
	modelContextSessionIdentifier string,
) error {
	result := s.database.Model(session.Stub()).Where(
		"callsign = ?",
		name,
	).Updates(
		map[string]any{
			"model_context_session": modelContextSessionIdentifier,
		},
	)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: %s", constant.ErrorCallsignNotFound, name)
	}

	return nil
}
