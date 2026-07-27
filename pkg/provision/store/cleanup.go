package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"time"
)

func (s *Store) Cleanup() {
	errors.PanicOnError(
		s.mapper.
			Table(s.tableName).
			Where(
				"created_at < ?",
				time.Now().Add(-constant.StoreRetentionAge),
			).
			Delete(&Run{}).Error,
	)
}
