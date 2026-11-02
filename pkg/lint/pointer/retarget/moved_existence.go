package retarget

import "github.com/funtimecoding/soil/pkg/lint/types/path_move"

func movedExistence(
	exists func(string) bool,
	moves []*path_move.Move,
) func(string) bool {
	return func(p string) bool {
		if exists(p) {
			return true
		}

		moved, found := apply(moves, p)

		return found && exists(moved)
	}
}
