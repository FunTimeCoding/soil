package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"path"
	"strings"
)

func Mentions(
	document string,
	content string,
	moves []*path_move.Move,
) bool {
	for _, m := range moves {
		segment := path.Base(m.From)

		if strings.Contains(content, segment) ||
			strings.Contains(document, segment) {
			return true
		}
	}

	return false
}
