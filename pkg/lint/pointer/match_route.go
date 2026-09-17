package pointer

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func MatchRoute(
	route string,
	paths []string,
) bool {
	want := strings.Split(strings.Trim(route, constant.Slash), constant.Slash)

	for _, p := range paths {
		have := strings.Split(strings.Trim(p, constant.Slash), constant.Slash)

		if matchSegments(want, have) {
			return true
		}
	}

	return false
}
