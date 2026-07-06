package in_range

import "github.com/funtimecoding/soil/pkg/math/ranges"

func Open(
	value float64,
	interval ranges.Range,
) bool {
	if value > interval.L && value < interval.R {
		return true
	}

	return false
}
