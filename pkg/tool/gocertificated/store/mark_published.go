package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"time"
)

func (s *Store) MarkPublished(
	serial []string,
	moment time.Time,
) error {
	if len(serial) == 0 {
		return nil
	}

	return s.database.
		Model(record.Stub()).
		Where("serial IN ?", serial).
		Update("published_at", moment).Error
}
