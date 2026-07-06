package chunk

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func splitAtBoundaries(
	content string,
	boundaries []int,
) []Chunk {
	var result []Chunk
	start := 0

	for _, boundary := range boundaries {
		if boundary <= start {
			continue
		}

		text := content[start:boundary]

		if len(text) > constant.ChunkSize {
			result = append(result, chunkMarkdown(text)...)
		} else {
			result = append(
				result,
				Chunk{
					Text:     text,
					Position: start,
				},
			)
		}

		start = boundary
	}

	if start < len(content) {
		text := content[start:]

		if len(text) > constant.ChunkSize {
			for _, c := range chunkMarkdown(text) {
				c.Position += start
				result = append(result, c)
			}
		} else {
			result = append(
				result,
				Chunk{
					Text:     text,
					Position: start,
				},
			)
		}
	}

	return result
}
