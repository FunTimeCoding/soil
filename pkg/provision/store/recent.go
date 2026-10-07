package store

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"time"
)

func (s *Store) Recent(limit int) ([]run.Run, error) {
	var result []run.Run

	return result, s.mapper.
		Table(s.tableName).
		Where("created_at > ?", time.Now().Add(-constant.StoreRetentionAge)).
		Order("created_at desc").
		Limit(limit).
		Find(&result).Error
}
