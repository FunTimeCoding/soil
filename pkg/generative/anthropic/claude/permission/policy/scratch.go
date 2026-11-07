package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"path/filepath"
	"strings"
)

func (w *Walk) scratch(word *syntax.Word) bool {
	path, okay := w.resolve(word)

	if !okay || path == "" ||
		strings.ContainsAny(path, constant.UnsafePathCharacter) {
		return false
	}

	if i := strings.IndexAny(path, constant.GlobCharacters); i >= 0 {
		directory := path[:strings.LastIndex(path[:i], "/")+1]

		for _, segment := range strings.Split(path[len(directory):], "/") {
			if strings.HasPrefix(segment, constant.HiddenSegmentPrefix) {
				return false
			}
		}

		path = directory
	}

	if !filepath.IsAbs(path) {
		if w.base == "" {
			return false
		}

		path = filepath.Join(w.base, path)
	}

	settled, okay := settle(path)

	return okay && w.inside(settled)
}
