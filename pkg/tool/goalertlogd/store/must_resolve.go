package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (s *Store) MustResolve(
	fingerprint string,
	at time.Time,
) {
	errors.PanicOnError(s.Resolve(fingerprint, at))
}
