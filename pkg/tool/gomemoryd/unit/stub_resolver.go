package unit

import (
	"github.com/funtimecoding/soil/pkg/lint/pointer/resolver"
	"github.com/funtimecoding/soil/pkg/markup/heading"
	"slices"
)

func stubResolver() *resolver.Resolver {
	result := resolver.New()
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
