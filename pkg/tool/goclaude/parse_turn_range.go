package goclaude

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strconv"
)

func parseTurnRange(s string) (int, int, bool) {
	parts := split.Dash(s)

	if len(parts) > 2 {
		return 0, 0, false
	}

	from, e := strconv.Atoi(parts[0])

	if e != nil || from < 1 {
		return 0, 0, false
	}

	to := from

	if len(parts) == 2 {
		to, e = strconv.Atoi(parts[1])

		if e != nil || to < from {
			return 0, 0, false
		}
	}

	return from, to, true
}
