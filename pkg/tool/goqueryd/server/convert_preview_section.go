package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/preview_section"
)

func convertPreviewSection(x *preview_section.Section) server.PreviewSection {
	blocks := make([]server.PreviewBlock, len(x.Blocks))

	for i, b := range x.Blocks {
		blocks[i] = server.PreviewBlock{
			Kind:      b.Kind,
			FirstLine: b.FirstLine,
			LastLine:  b.LastLine,
			Tokens:    b.Tokens,
		}
	}

	return server.PreviewSection{
		Level:     x.Level,
		Title:     x.Title,
		FirstLine: x.FirstLine,
		LastLine:  x.LastLine,
		Tokens:    x.Tokens,
		Blocks:    blocks,
	}
}
