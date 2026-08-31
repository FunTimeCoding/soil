package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/fable_snapshot"
)

func (s *Store) SaveFableSnapshot(
	percent int,
	reset string,
) {
	errors.PanicOnError(
		s.database.Create(fable_snapshot.New(percent, reset, s.clock())).Error,
	)
}
