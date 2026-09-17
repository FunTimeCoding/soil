package pointer_tester

import "github.com/funtimecoding/soil/pkg/lint"

func Checker(existing ...string) lint.Checker {
	return lint.Pointers(Resolver(existing...), discard)
}
