package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"time"
)

func (s *Store) Revoke(
	serial string,
	moment time.Time,
) error {
	return s.database.
		Model(record.Stub()).
		Where("serial = ? AND revoked_at IS NULL", serial).
		Update("revoked_at", moment).Error
}
