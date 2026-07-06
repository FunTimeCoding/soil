package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) ClearBindings() {
	errors.PanicOnError(
		s.database.Model(session.Stub()).
			Where("model_context_session != ''").
			Update("model_context_session", "").Error,
	)
}
