package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) UnenrichedSessions() []session.Session {
	var result []session.Session
	errors.PanicOnError(
		s.database.
			Where("slug = '' OR slug IS NULL OR callsign IS NOT NULL OR started_at IS NULL OR started_at = ''").
			Find(&result).Error,
	)

	return result
}
