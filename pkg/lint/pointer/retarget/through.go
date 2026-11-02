package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
)

func Through(
	r *resolver.Resolver,
	moves []*path_move.Move,
) *resolver.Resolver {
	result := *r
	result.Exists = movedExistence(r.Exists, moves)
	result.SiblingExists = movedExistence(r.SiblingExists, moves)

	return &result
}
