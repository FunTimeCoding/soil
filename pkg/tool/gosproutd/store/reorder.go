package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
)

func (s *Store) Reorder(identifiers []uint) {
	for i, v := range identifiers {
		errors.PanicOnError(
			s.mapper.Model(seed.Stub()).
				Where("identifier = ?", v).
				Update("position", i+1).Error,
		)
	}
}
