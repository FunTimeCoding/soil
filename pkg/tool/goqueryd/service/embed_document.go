package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"
	"time"
)

func (s *Service) embedDocument(
	x context.Context,
	p record.PendingDocument,
) (int, error) {
	chunks := chunk.Document(p.Body, p.Path, s.reranker)
	texts := make([]string, len(chunks))

	for i, c := range chunks {
		texts[i] = c.Text
	}

	embeddings, e := s.embedTexts(x, texts)

	if e != nil {
		return 0, e
	}

	now := time.Now().UTC().Format(time.RFC3339)

	for i, embedding := range embeddings {
		s.store.InsertEmbedding(p.Hash, i, chunks[i].Position, embedding, now)
	}

	return len(chunks), nil
}
