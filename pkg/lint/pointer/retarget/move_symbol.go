package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func moveSymbol(
	moves *move_set.Set,
	directory string,
	tail string,
) (string, bool) {
	name, qualified := strings.CutPrefix(tail, constant.Dot)

	if !qualified {
		return "", false
	}

	for _, m := range moves.Symbols {
		if m.Directory != directory || !strings.HasPrefix(name, m.Name) {
			continue
		}

		rest := name[len(m.Name):]

		if rest != "" && !strings.HasPrefix(rest, constant.Dot) &&
			!strings.HasPrefix(rest, "(") {
			continue
		}

		return join.Empty(m.NewDirectory, constant.Dot, m.NewName, rest), true
	}

	return "", false
}
