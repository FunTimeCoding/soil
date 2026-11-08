package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"go/types"
)

func moment(v *types.Var) bool {
	parts := segment.Segments(v.Name())

	return parts[len(parts)-1] == constant.MomentSegment || isTime(v.Type())
}
