package git

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

func PatternMatcher(patterns []gitignore.Pattern) func(string) bool {
	m := gitignore.NewMatcher(patterns)

	return func(p string) bool {
		parts := split.Slash(p)

		return m.Match(parts, false) || m.Match(parts, true)
	}
}
