package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"time"
)

func (s *Store) ByTimeRange(
	start time.Time,
	end time.Time,
) ([]record.Record, error) {
	var result []record.Record

	return result, s.database.
		Where("started_at >= ? AND started_at < ?", start, end).
		Order("started_at").
		Find(&result).Error
}
