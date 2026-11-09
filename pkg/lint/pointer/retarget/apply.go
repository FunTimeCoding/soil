package retarget

import (
	"github.com/funtimecoding/soil/pkg/lint/types/path_move"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func apply(
	moves []*path_move.Move,
	target string,
) (string, bool) {
	var best *path_move.Move

	for _, m := range moves {
		if target != m.Old &&
			!strings.HasPrefix(target, join.Empty(m.Old, stringsConstant.Slash)) {
			continue
		}

		if best == nil || len(m.Old) > len(best.Old) {
			best = m
		}
	}

	if best == nil {
		return target, false
	}

	return join.Empty(best.New, target[len(best.Old):]), true
}
