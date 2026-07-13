package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"time"
)

func (s *Store) Prune(cutoff time.Time) (int, error) {
	r := s.database.
		Where("started_at < ?", cutoff).
		Delete(record.Stub())

	return int(r.RowsAffected), r.Error
}
