package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
)

func Gitignored(ignored ...string) lint.Checker {
	r := resolver.New()
	r.Roots = Roots()
	r.Ignored = exists(ignored)

	return lint.Pointers(r, discard)
}
