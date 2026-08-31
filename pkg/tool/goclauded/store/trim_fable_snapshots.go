package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/fable_snapshot"
)

func (s *Store) TrimFableSnapshots() {
	errors.PanicOnError(
		s.database.Where("created_at < ?", s.clock().AddDate(0, 0, -7)).Delete(
			fable_snapshot.Stub(),
		).Error,
	)
}
