package scale

import (
	"github.com/funtimecoding/soil/pkg/math/normalize"
	"github.com/funtimecoding/soil/pkg/math/numeric"
)

func Value[T numeric.Number](
	from T,
	to T,
	factor float64,
) T {
	normalize.Float(&factor, 0, 1)

	return from + T(float64(to-from)*factor)
}
