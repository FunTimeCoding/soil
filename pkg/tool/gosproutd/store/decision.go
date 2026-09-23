package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"gorm.io/gorm"
)

func (s *Store) Decision(identifier uint) *decision.Decision {
	result := decision.Stub()
	errors.PanicOnError(
		s.mapper.Preload(
			"Choices",
			func(d *gorm.DB) *gorm.DB { return d.Order("position") },
		).Preload(
			"Frames",
		).Preload(
			"Turns",
			func(d *gorm.DB) *gorm.DB {
				return d.Order(constant.IdentifierColumn)
			},
		).First(result, identifier).Error,
	)

	return result
}
