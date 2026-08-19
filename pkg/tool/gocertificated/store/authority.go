package store

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"

func (s *Store) Authority(name string) (*record.Record, error) {
	var result []record.Record
	e := s.database.
		Where("name = ? AND revoked_at IS NULL", name).
		Limit(1).
		Find(&result).Error

	if e != nil || len(result) == 0 {
		return nil, e
	}

	return &result[0], nil
}
