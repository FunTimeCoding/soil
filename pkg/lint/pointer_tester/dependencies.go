package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Dependencies(modules ...string) lint.Checker {
	r := resolver.New()
	r.Roots = Roots()
	r.Dependency = func(span string) bool {
		for _, m := range modules {
			_, rest, _ := strings.Cut(m, stringsConstant.Slash)

			for _, form := range []string{m, rest} {
				if span == form ||
					strings.HasPrefix(span, join.Empty(form, stringsConstant.Slash)) {
					return true
				}
			}
		}

		return false
	}

	return lint.Pointers(r, discard)
}
