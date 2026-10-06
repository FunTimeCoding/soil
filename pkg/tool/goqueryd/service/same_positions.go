package service

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"

func samePositions(
	stored map[int]int,
	chunks []chunk.Chunk,
) bool {
	if len(stored) != len(chunks) {
		return false
	}

	for i, c := range chunks {
		if position, okay := stored[i]; !okay || position != c.Position {
			return false
		}
	}

	return true
}
