package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"path"
	"strings"
)

func Mentions(
	document string,
	content string,
	moves *move_set.Set,
) bool {
	var segments []string

	for _, m := range moves.Paths {
		segments = append(segments, path.Base(m.From))
	}

	for _, m := range moves.Symbols {
		segments = append(segments, path.Base(m.Directory))
	}

	for _, segment := range segments {
		if strings.Contains(content, segment) ||
			strings.Contains(document, segment) {
			return true
		}
	}

	return false
}
