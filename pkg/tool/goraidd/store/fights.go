package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/raid"
)

func (s *Store) Fights() []raid.Fight {
	var fights []raid.Fight
	errors.PanicOnError(s.mapper.Order("timestamp desc").Find(&fights).Error)

	return fights
}
