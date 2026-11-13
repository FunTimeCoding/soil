package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

func IgnoreMatcher(path string) func(string) bool {
	patterns, e := gitignore.ReadPatterns(osfs.New(path), nil)
	errors.PanicOnError(e)

	return PatternMatcher(patterns)
}
