package lint

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func dependencyMatcher(modules []string) func(string) bool {
	return func(span string) bool {
		for _, m := range modules {
			_, rest, _ := strings.Cut(m, constant.Slash)

			for _, form := range []string{m, rest} {
				if form == "" {
					continue
				}

				if span == form ||
					strings.HasPrefix(span, join.Empty(form, constant.Slash)) {
					return true
				}
			}
		}

		return false
	}
}
