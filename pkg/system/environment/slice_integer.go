package environment

import "github.com/funtimecoding/soil/pkg/strings"

func SliceInteger(name string) []int {
	return strings.MustToIntegers(Slice(name))
}
