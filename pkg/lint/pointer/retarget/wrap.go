package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func wrap(
	link bool,
	span string,
) string {
	if link {
		return join.Empty("(", span, ")")
	}

	tick := string(constant.Backtick)

	return join.Empty(tick, span, tick)
}
