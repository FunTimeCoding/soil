package chunk

import "github.com/funtimecoding/soil/pkg/face"

func assemble(
	content string,
	tables []*Table,
	pulled map[int]bool,
	c face.TokenCounter,
) []Chunk {
	if len(pulled) == 0 {
		return segment(content, 0)
	}

	var result []Chunk
	start := 0

	for i, t := range tables {
		if !pulled[i] {
			continue
		}

		result = append(result, prose(content[start:t.start], start)...)
		result = append(result, splitTable(content, t, c)...)
		start = t.end
	}

	return append(result, prose(content[start:], start)...)
}
