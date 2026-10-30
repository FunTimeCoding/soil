package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"time"
)

func (s *Service) embedEach(
	x context.Context,
	hash string,
	chunks []chunk.Chunk,
) (int, error) {
	embeddings := make([][]float32, len(chunks))
	refusals := make([]error, len(chunks))
	var refused error

	for i, c := range chunks {
		v, e := s.embedTexts(x, []string{c.Text})

		if e == nil {
			embeddings[i] = v[0]

			continue
		}

		if !validation.Is(e) {
			return 0, e
		}

		refusals[i] = e

		if refused == nil {
			refused = e
		}
	}

	now := time.Now().UTC().Format(time.RFC3339)
	result := 0

	for i, c := range chunks {
		if refusals[i] != nil {
			s.store.InsertEmbeddingFailure(
				hash,
				i,
				c.Position,
				refusals[i].Error(),
				now,
			)

			continue
		}

		s.store.InsertEmbedding(hash, i, c.Position, embeddings[i], now)
		result++
	}

	return result, refused
}
