package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/constant"
)

func (s *Store) AbandonRunning() {
	errors.PanicOnError(
		s.mapper.
			Table(s.tableName).
			Where("status = ?", constant.StoreStatusRunning).
			Updates(
				map[string]any{
					"status":       constant.StoreStatusError,
					"error_output": constant.StoreAbandoned,
				},
			).Error,
	)
}
