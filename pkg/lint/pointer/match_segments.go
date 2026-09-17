package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func matchSegments(
	want []string,
	have []string,
) bool {
	for i, segment := range want {
		if segment == constant.RouteEllipsis {
			return true
		}

		if i >= len(have) {
			return false
		}

		if segment == have[i] ||
			strings.HasPrefix(segment, "{") ||
			strings.HasPrefix(have[i], "{") {
			continue
		}

		return false
	}

	return len(want) == len(have)
}
