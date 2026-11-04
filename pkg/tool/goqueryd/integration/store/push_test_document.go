//go:build local

package store

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
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
	if e := s.EnsurePushCollection(collection); e != nil {
		return e
	}

	now := "2024-01-01T00:00:00Z"
	hash := store.HashContent(body)
	title := store.ExtractTitle(body, path)

	if e := s.InsertContent(hash, body, now); e != nil {
		return e
	}

	if e := s.InsertDocument(collection, path, title, hash, now); e != nil {
		return e
	}

	r := mock_reranker.New()
	defer errors.LogClose(r)
	chunks := chunk.Document(body, path, r)
	texts := make([]string, len(chunks))

	for i, c := range chunks {
		texts[i] = c.Text
	}

	embeddings, e := o.Embed(context.Background(), texts)

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
