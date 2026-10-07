package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
)

func (s *Store) Create(r *run.Run) {
	errors.PanicOnError(s.mapper.Table(s.tableName).Create(r).Error)
}
