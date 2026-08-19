package store

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (s *Store) Unpublished() ([]record.Record, error) {
	var result []record.Record

	return result, s.database.
		Where("name != ? AND published_at IS NULL AND revoked_at IS NULL", "").
		Order("not_before").
		Find(&result).Error
}
