package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
)

func (s *Store) MustCreateEvent(v event.Event) {
	errors.PanicOnError(s.CreateEvent(v))
}
