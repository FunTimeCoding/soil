package store

import (
	"github.com/bmatcuk/doublestar/v4"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func globFiles(
	root string,
	pattern string,
) []string {
	var result []string
	base := os.DirFS(root)
	matches, e := doublestar.Glob(base, pattern)
	errors.PanicOnError(e)

	for _, match := range matches {
		if isHidden(match) {
			continue
		}

		result = append(result, match)
	}

	return result
}
