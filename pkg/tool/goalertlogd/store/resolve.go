package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"time"
)

func (s *Store) Resolve(fingerprint string) error {
	return s.database.
		Model(record.Stub()).
		Where("fingerprint = ? AND ended_at IS NULL", fingerprint).
		Update("ended_at", time.Now()).Error
}
