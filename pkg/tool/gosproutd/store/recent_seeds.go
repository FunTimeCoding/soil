package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
)

func (s *Store) RecentSeeds() []seed.Seed {
	var result []seed.Seed
	errors.PanicOnError(s.mapper.Order("modified_at DESC").Find(&result).Error)

	return result
}
