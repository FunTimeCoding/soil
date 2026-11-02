package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func symbolSpan(
	moves *move_set.Set,
	rest string,
) string {
	directory := pointer.PackageDirectory(rest)

	if directory == "" || !strings.HasPrefix(rest, directory) {
		return ""
	}

	tail := rest[len(directory):]
	moved, found := moveSymbol(moves, directory, tail)

	if !found {
		moved, found = apply(moves.Paths, directory)

		if !found {
			return ""
		}

		moved = join.Empty(moved, tail)
	}

	return join.Empty(constant.SchemeGo, moved)
}
