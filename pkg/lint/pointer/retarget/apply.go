package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func apply(
	moves []*path_move.Move,
	target string,
) (string, bool) {
	var best *path_move.Move

	for _, m := range moves {
		if target != m.From &&
			!strings.HasPrefix(target, join.Empty(m.From, constant.Slash)) {
			continue
		}

		if best == nil || len(m.From) > len(best.From) {
			best = m
		}
	}

	if best == nil {
		return target, false
	}

	return join.Empty(best.To, target[len(best.From):]), true
}
