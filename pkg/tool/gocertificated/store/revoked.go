package store

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (s *Store) Revoked() ([]record.Record, error) {
	var result []record.Record

	return result, s.database.
		Where("revoked_at IS NOT NULL").
		Order("revoked_at").
		Find(&result).Error
}
