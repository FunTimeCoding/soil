package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
)

func (s *Store) Fights() []fight.Fight {
	var fights []fight.Fight
	errors.PanicOnError(s.mapper.Order("timestamp desc").Find(&fights).Error)

	return fights
}
