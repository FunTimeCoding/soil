package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/rate_snapshot"
)

func (s *Store) TrimRateSnapshots() {
	errors.PanicOnError(
		s.database.Where(
			"created_at < ?",
			s.clock().AddDate(0, 0, -7),
		).Delete(rate_snapshot.Stub()).Error,
	)
}
