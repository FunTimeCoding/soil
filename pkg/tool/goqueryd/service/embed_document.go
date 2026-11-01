package service

import (
	"context"
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
	refusals := s.overWindow(chunks)
	embeddings, e := s.embedFitting(x, chunks, refusals)

	if e != nil {
		return 0, e
	}

	now := time.Now().UTC().Format(time.RFC3339)
	result := 0
	var refused error

	for i, c := range chunks {
		if refusals[i] != nil {
			s.store.InsertEmbeddingFailure(
				hash,
				i,
				c.Position,
				refusals[i].Error(),
				now,
			)

			if refused == nil {
				refused = refusals[i]
			}

			continue
		}

		s.store.InsertEmbedding(hash, i, c.Position, embeddings[i], now)
		result++
	}

	return result, refused
}
