package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (s *Store) Cleanup() {
	errors.PanicOnError(
		s.mapper.
			Table(s.tableName).
			Where("created_at < ?", time.Now().Add(-RetentionAge)).
			Delete(&Run{}).Error,
	)
}
