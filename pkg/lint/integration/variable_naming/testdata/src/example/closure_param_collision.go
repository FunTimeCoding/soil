package example

import (
	"sort"
	"strconv"
)

func ClosureParamCollision() {
	values := []string{"3", "1", "2"}
	sort.Slice(
		values,
		func(a, b int) bool {
			x, xFail := strconv.Atoi(values[a])
			y, yFail := strconv.Atoi(values[b])

			if xFail != nil || yFail != nil {
				return values[a] > values[b]
			}

			return x > y
		},
	)
}
