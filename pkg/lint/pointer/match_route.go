package pointer

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func MatchRoute(
	route string,
	paths []string,
) bool {
	want := strings.Split(
		strings.Trim(route, stringsConstant.Slash),
		stringsConstant.Slash,
	)

	for _, p := range paths {
		have := strings.Split(
			strings.Trim(p, stringsConstant.Slash),
			stringsConstant.Slash,
		)

		if MatchSegments(want, have) {
			return true
		}
	}

	return false
}
