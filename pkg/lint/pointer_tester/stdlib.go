package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
)

func Stdlib(packages ...string) lint.Checker {
	r := resolver.New()
	r.Roots = Roots()
	r.Stdlib = exists(packages)

	return lint.Pointers(r, discard)
}
