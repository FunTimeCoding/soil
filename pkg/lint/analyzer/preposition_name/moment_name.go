package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/segment"
)

func momentName(name string) bool {
	parts := segment.Segments(name)
	last := len(parts) - 1

	return parts[last] == constant.MomentSegment ||
		last > 0 &&
			parts[last-1] == constant.MomentSegment &&
			constant.MomentSuffix[parts[last]]
}
