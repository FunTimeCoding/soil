package normalize_change

import "github.com/funtimecoding/soil/pkg/math/numeric"

func Value[T numeric.Number](
	now T,
	change T,
	minimum T,
	maximum T,
) T {
	if maximum > minimum && now+change > maximum {
		return maximum - now
	}

	if now+change < minimum {
		return (now - minimum) * -1
	}

	return change
}
