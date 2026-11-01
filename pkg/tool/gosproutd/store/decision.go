package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"gorm.io/gorm"
)

func (s *Store) Decision(identifier uint) (*decision.Decision, error) {
	result := decision.Stub()
	e := s.mapper.Preload(
		"Choices",
		func(d *gorm.DB) *gorm.DB { return d.Order("position") },
	).Preload(
		"Frames",
	).Preload(
		"Turns",
		func(d *gorm.DB) *gorm.DB {
			return d.Order(constant.IdentifierColumn)
		},
	).First(result, identifier).Error

	if errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, not_found.New(constant.DecisionKind, identifier)
	}

	errors.PanicOnError(e)

	return result, nil
}
