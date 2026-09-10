package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) BoundSessions() []session.Session {
	var result []session.Session
	errors.PanicOnError(
		s.database.Where("model_context_session != ''").Find(&result).Error,
	)

	return result
}
