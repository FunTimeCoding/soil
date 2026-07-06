package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) CreateDiscovered(identifier string) {
	errors.PanicOnError(
		s.database.Create(session.NewDiscovered(identifier, s.clock())).Error,
	)
}
