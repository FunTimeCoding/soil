package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"time"
)

func (s *Service) embedDocument(
	x context.Context,
	hash string,
	path string,
	body string,
) (int, error) {
	chunks := chunk.Document(body, path, s.reranker)
	texts := make([]string, len(chunks))

	for i, c := range chunks {
		texts[i] = c.Text
	}

	embeddings, e := s.embedTexts(x, texts)

	if e != nil {
		if !validation.Is(e) {
			return 0, e
		}

		return s.embedEach(x, hash, chunks)
	}

	now := time.Now().UTC().Format(time.RFC3339)

	for i, embedding := range embeddings {
		s.store.InsertEmbedding(hash, i, chunks[i].Position, embedding, now)
	}

	return len(chunks), nil
}
