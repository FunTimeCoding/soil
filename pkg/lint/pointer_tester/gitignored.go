package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func Gitignored(ignored ...string) lint.Checker {
	r := pointer.New()
	r.Roots = Roots()
	r.Ignored = exists(ignored)

	return lint.Pointers(r, discard)
}
