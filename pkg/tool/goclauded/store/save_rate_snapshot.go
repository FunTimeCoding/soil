package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/rate_snapshot"
	"time"
)

func (s *Store) SaveRateSnapshot(
	fiveHourPercent int,
	sevenDayPercent int,
	fiveHourReset time.Time,
	sevenDayReset time.Time,
) {
	errors.PanicOnError(
		s.database.Create(
			rate_snapshot.New(
				fiveHourPercent,
				sevenDayPercent,
				fiveHourReset,
				sevenDayReset,
				s.clock(),
			),
		).Error,
	)
}
