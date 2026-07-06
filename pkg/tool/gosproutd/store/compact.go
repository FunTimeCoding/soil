package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
)

func (s *Store) Compact() {
	var result []seed.Seed
	errors.PanicOnError(s.mapper.Order("position").Find(&result).Error)

	for i, v := range result {
		target := i + 1

		if v.Position == target {
			continue
		}

		errors.PanicOnError(
			s.mapper.Model(seed.Stub()).
				Where("identifier = ?", v.Identifier).
				Update("position", target).Error,
		)
	}
}
