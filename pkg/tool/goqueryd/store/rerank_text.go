package store

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
)

func RerankText(
	body string,
	r *search.Result,
	t face.TokenCounter,
) string {
	if body == "" {
		return r.Snippet
	}

	chunks := chunk.Document(body, r.Path, t)

	if r.ChunkPosition >= 0 {
		for _, c := range chunks {
			if c.Position == r.ChunkPosition {
				return c.Text
			}
		}
	}

	offset := LineOffset(body, r.SnippetLine)

	for _, c := range chunks {
		if offset < c.Position+c.Length {
			return c.Text
		}
	}

	return chunks[len(chunks)-1].Text
}
