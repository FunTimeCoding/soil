package chunk

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func Document(
	content string,
	filename string,
) []Chunk {
	if len(content) <= constant.ChunkSize {
		return []Chunk{{Text: content, Position: 0}}
	}

	if strings.HasSuffix(filename, library.GoExtension) {
		result := chunkGoSource(content)

		if len(result) > 0 {
			return result
		}
	}

	return chunkMarkdown(content)
}
