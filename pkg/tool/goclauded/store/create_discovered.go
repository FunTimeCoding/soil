package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Store) CreateDiscovered(identifier string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, found, e := s.FindSession(identifier)
	errors.PanicOnError(e)

	if found {
		return
	}

	errors.PanicOnError(
		s.database.Create(session.NewDiscovered(identifier, s.clock())).Error,
	)
}
