package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"time"
)

func (s *Store) Expiring(before time.Time) ([]record.Record, error) {
	var result []record.Record

	return result, s.database.
		Where("not_after < ? AND revoked_at IS NULL", before).
		Order("not_after").
		Find(&result).Error
}
