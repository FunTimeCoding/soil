package pointer_tester

import "github.com/funtimecoding/soil/pkg/lint"

func Implicit(
	bases []string,
	existing ...string,
) lint.Checker {
	r := Resolver(existing...)
	r.ImplicitBases = bases

	return lint.Pointers(r, discard)
}
