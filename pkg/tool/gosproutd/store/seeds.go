package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
)

func (s *Store) Seeds() []seed.Seed {
	var result []seed.Seed
	errors.PanicOnError(s.mapper.Order("position").Find(&result).Error)

	return result
}
