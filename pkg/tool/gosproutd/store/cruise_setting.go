package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/cruise"
)

func (s *Store) CruiseSetting(session string) *cruise.Cruise {
	result := cruise.Stub()
	found := s.mapper.Where("session = ?", session).Limit(1).Find(result)
	errors.PanicOnError(found.Error)

	if found.RowsAffected == 0 {
		return cruise.New(session, constant.CruiseOff, 0)
	}

	return result
}
