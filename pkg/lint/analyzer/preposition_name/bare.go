package preposition_name

import "github.com/funtimecoding/soil/pkg/lint/segment"

func bare(name string) bool {
	parts := segment.Segments(name)

	return len(parts) == 1 && placedWord(parts, 0) != ""
}
