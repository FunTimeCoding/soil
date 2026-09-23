//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
)

func embedTestDocuments(
	s *store.Store,
	o *ollama.Client,
) error {
	pending := s.PendingEmbeddings()
	now := "2024-01-01T00:00:00Z"

	for _, p := range pending {
		chunks := chunk.Document(p.Body, p.Path)
		texts := make([]string, len(chunks))

		for i, c := range chunks {
			texts[i] = c.Text
		}

		embeddings, e := o.Embed(texts)

		if e != nil {
			return e
		}

		for i, embedding := range embeddings {
			s.InsertEmbedding(p.Hash, i, chunks[i].Position, embedding, now)
		}
	}

	return nil
}
