package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
	"strings"
)

func IgnoreMatcher(path string) func(string) bool {
	patterns, e := gitignore.ReadPatterns(osfs.New(path), nil)
	errors.PanicOnError(e)
	m := gitignore.NewMatcher(patterns)

	return func(p string) bool {
		parts := strings.Split(p, "/")

		return m.Match(parts, false) || m.Match(parts, true)
	}
}
