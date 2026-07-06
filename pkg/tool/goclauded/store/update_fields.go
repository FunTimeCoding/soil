package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) UpdateFields(
	identifier string,
	updates map[string]any,
) {
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("identifier = ?", identifier).
			Updates(updates).Error,
	)
}
