//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
)

func pushTestDocument(
	s *store.Store,
	o *ollama.Client,
	collection string,
	path string,
	body string,
	sourceType string,
) error {
	s.EnsurePushCollection(collection)
	now := "2024-01-01T00:00:00Z"
	hash := store.HashContent(body)
	title := store.ExtractTitle(body, path)
	s.InsertContent(hash, body, now)
	s.InsertDocument(collection, path, title, hash, now)
	chunks := chunk.Document(body, path)
	texts := make([]string, len(chunks))

	for i, c := range chunks {
		texts[i] = c.Text
	}

	embeddings, e := o.Embed(texts)

	if e != nil {
		return e
	}

	for i, embedding := range embeddings {
		s.InsertEmbedding(hash, i, chunks[i].Position, embedding, now)
	}

	if sourceType != "" {
		s.SetSourceType(collection, path, sourceType)
	}

	return nil
}
