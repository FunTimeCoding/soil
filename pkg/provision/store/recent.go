package store

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"time"
)

func (s *Store) Recent(limit int) ([]Run, error) {
	var result []Run

	return result, s.mapper.
		Table(s.tableName).
		Where("created_at > ?", time.Now().Add(-constant.StoreRetentionAge)).
		Order("created_at desc").
		Limit(limit).
		Find(&result).Error
}
