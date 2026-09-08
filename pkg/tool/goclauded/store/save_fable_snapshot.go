package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/fable_snapshot"
	"time"
)

func (s *Store) SaveFableSnapshot(
	percent int,
	reset string,
	resetAt *time.Time,
) {
	errors.PanicOnError(
		s.database.Create(
			fable_snapshot.New(percent, reset, resetAt, s.clock()),
		).Error,
	)
}
