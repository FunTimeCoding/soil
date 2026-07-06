package strings

import "github.com/funtimecoding/soil/pkg/math/normalize"

func Intensity(
	v []string,
	intensity float64,
) string {
	normalize.Float(&intensity, 0, 1)
	length := len(v)
	index := int(float64(length) * intensity)

	if index == length {
		index--
	}

	return v[index]
}
