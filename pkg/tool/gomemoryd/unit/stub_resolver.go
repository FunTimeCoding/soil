package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/markup/heading"
	"slices"
)

func stubResolver() *pointer.Resolver {
	result := pointer.New()
	result.Roots = []string{"doc", "pkg"}
	result.Exists = func(p string) bool {
		return slices.Contains(
			[]string{"doc", "doc/real.md", "pkg", "pkg/real.go"},
			p,
		)
	}
	result.Headings = func(p string) ([]*heading.Heading, bool) {
		if p != "doc/real.md" {
			return nil, false
		}

		return heading.Parse("# Real\n\n## Usage\n"), true
	}

	return result
}
