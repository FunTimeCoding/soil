package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
)

func Declaring(
	declared []string,
	existing ...string,
) lint.Checker {
	r := Resolver(existing...)
	r.Declares = func(
		directory string,
		name string,
	) bool {
		return slices.Contains(
			declared,
			join.Empty(directory, stringsConstant.Dot, name),
		)
	}

	return lint.Pointers(r, discard)
}
