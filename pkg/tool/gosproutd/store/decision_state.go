package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func (s *Store) DecisionState(identifier uint) constant.State {
	var result string
	errors.PanicOnError(
		s.mapper.Model(decision.Stub()).Where(
			"identifier = ?",
			identifier,
		).Select(constant.StateColumn).Row().Scan(&result),
	)

	return constant.State(result)
}
