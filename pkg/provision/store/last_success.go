package store

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"time"
)

func (s *Store) LastSuccess() (time.Time, error) {
	var result Run
	e := s.mapper.
		Table(s.tableName).
		Where("status = ?", constant.StoreStatusSuccess).
		Order("created_at desc").
		First(&result).Error

	return result.CreatedAt, e
}
