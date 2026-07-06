package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/usage_snapshot"
)

func (s *Store) TrimUsageSnapshots() {
	errors.PanicOnError(
		s.database.Where(
			"created_at < ?",
			s.clock().AddDate(0, 0, -7),
		).Delete(usage_snapshot.Stub()).Error,
	)
}
