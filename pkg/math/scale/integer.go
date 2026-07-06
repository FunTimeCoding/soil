package scale

import "github.com/funtimecoding/soil/pkg/math/normalize"

// Integer
// At factor 0: from value
// At factor 1: to value
func Integer(
	from int,
	to int,
	factor float64,
) int {
	normalize.Float(&factor, 0, 1)

	return from + int(float64(to-from)*factor)
}
