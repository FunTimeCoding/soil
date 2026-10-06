package chunk

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func Document(
	content string,
	filename string,
	c face.TokenCounter,
) []Chunk {
	if !strings.HasSuffix(filename, library.GoExtension) {
		return chunkText(content, c)
	}

	if len(content) <= constant.ChunkSize {
		return []Chunk{{Text: content, Position: 0, Length: len(content)}}
	}

	if result := chunkGoSource(content); len(result) > 0 {
		return result
	}

	return chunkMarkdown(content)
}
