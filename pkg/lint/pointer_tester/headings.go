package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/markup/heading"
)

func Headings(
	contents map[string]string,
	existing ...string,
) lint.Checker {
	r := Resolver(existing...)
	r.Headings = func(path string) ([]*heading.Heading, bool) {
		content, found := contents[path]

		if !found {
			return nil, false
		}

		return heading.Parse(content), true
	}

	return lint.Pointers(r, discard)
}
