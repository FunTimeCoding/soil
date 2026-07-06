package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/usage_snapshot"
)

func (s *Store) SaveUsageSnapshot(
	sessionPercent int,
	weeklyPercent int,
	sonnetPercent int,
	creditPercent int,
) {
	errors.PanicOnError(
		s.database.Create(
			usage_snapshot.New(
				sessionPercent,
				weeklyPercent,
				sonnetPercent,
				creditPercent,
				s.clock(),
			),
		).Error,
	)
}
