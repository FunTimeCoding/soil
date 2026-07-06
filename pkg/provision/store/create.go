package store

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Store) Create(r *Run) {
	errors.PanicOnError(s.mapper.Table(s.tableName).Create(r).Error)
}
