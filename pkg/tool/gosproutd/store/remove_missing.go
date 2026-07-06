package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
)

func (s *Store) RemoveMissing(paths []string) {
	if len(paths) == 0 {
		errors.PanicOnError(s.mapper.Where("1 = 1").Delete(seed.Stub()).Error)

		return
	}

	errors.PanicOnError(
		s.mapper.Where("path NOT IN ?", paths).Delete(seed.Stub()).Error,
	)
}
