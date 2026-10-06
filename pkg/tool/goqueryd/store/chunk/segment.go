package chunk

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func segment(
	text string,
	offset int,
) []Chunk {
	if len(text) <= constant.ChunkSize {
		return []Chunk{{Text: text, Position: offset, Length: len(text)}}
	}

	result := chunkMarkdown(text)

	for i := range result {
		result[i].Position += offset
	}

	return result
}
