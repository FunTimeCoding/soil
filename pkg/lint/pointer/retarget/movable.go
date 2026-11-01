package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
)

func movable(
	r *resolver.Resolver,
	link bool,
	span string,
) bool {
	if link && !pointer.IsPath(span) {
		return true
	}

	switch pointer.Classify(span, r.Roots) {
	case constant.PointerClassRepository,
		constant.PointerClassSibling,
		constant.PointerClassShort:

		return true
	}

	return false
}
