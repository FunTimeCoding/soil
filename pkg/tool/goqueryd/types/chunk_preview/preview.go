package chunk_preview

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/preview_chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/preview_section"
)

type Preview struct {
	Model           string
	Window          int
	TokenLimit      int
	EmbedTokenLimit int
	Chunks          []*preview_chunk.Chunk
	Sections        []*preview_section.Section
}
