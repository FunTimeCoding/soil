package floats

import "github.com/funtimecoding/soil/pkg/math/round"

func ToStringRounded(number float64) string {
	return ToString(round.Round(number, 1))
}
