package layer

import (
	"github.com/funtimecoding/soil/pkg/source/types/layer_change"
	"maps"
	"slices"
	"strings"
)

func (l *Layer) Changes() []*layer_change.Change {
	moves := map[string]string{}

	for _, c := range l.stack() {
		maps.Copy(moves, c.moves)
	}

	arrived := map[string]bool{}

	for _, to := range moves {
		arrived[to] = true
	}

	var result []*layer_change.Change

	for _, path := range l.touched() {
		if arrived[path] {
			continue
		}

		destination := path

		for range len(moves) {
			to, moved := moves[destination]

			if !moved {
				break
			}

			destination = to
		}

		current, e := l.Read(destination)

		if e != nil {
			current = nil
		}

		result = append(
			result,
			layer_change.New(path, destination, l.originalOf(path), current),
		)
	}

	slices.SortFunc(
		result,
		func(a *layer_change.Change, b *layer_change.Change) int {
			return strings.Compare(a.Path, b.Path)
		},
	)

	return result
}
