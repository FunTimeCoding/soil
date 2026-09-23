package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) OpenDecisions(frameIdentifier uint) []*decision.Decision {
	var result []*decision.Decision
	errors.PanicOnError(
		s.mapper.Preload("Choices").Joins(
			"JOIN decision_frame ON decision_frame.decision_identifier = decision.identifier",
		).Where(
			"decision_frame.frame_identifier = ? AND decision.state = ? AND decision.resolution = ?",
			frameIdentifier,
			constant.StateOpen,
			"",
		).Order("decision.identifier").Find(&result).Error,
	)

	return result
}
