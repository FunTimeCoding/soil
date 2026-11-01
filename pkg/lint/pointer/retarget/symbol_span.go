package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func symbolSpan(
	moves []*path_move.Move,
	rest string,
) string {
	directory := pointer.PackageDirectory(rest)

	if directory == "" || !strings.HasPrefix(rest, directory) {
		return ""
	}

	moved, found := apply(moves, directory)

	if !found {
		return ""
	}

	return join.Empty(constant.SchemeGo, moved, rest[len(directory):])
}
