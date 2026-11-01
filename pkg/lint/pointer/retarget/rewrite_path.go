package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path"
	"path/filepath"
	"strings"
)

func rewritePath(
	moves []*path_move.Move,
	source string,
	normalized string,
	target string,
) string {
	moved, found := apply(moves, target)

	if !found {
		return ""
	}

	result := moved

	if strings.HasSuffix(target, normalized) {
		implicit := strings.TrimSuffix(target, normalized)

		if implicit != "" {
			head, _ := apply(
				moves,
				strings.TrimSuffix(implicit, constant.Slash),
			)
			prefix := join.Empty(head, constant.Slash)

			if strings.HasPrefix(moved, prefix) {
				result = strings.TrimPrefix(moved, prefix)
			}
		}
	} else if relative, e := filepath.Rel(path.Dir(source), moved); e == nil {
		result = filepath.ToSlash(relative)
	}

	if result == normalized {
		return ""
	}

	return result
}
