package service

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"time"
)

func (s *Service) PushDocument(
	x context.Context,
	collection string,
	path string,
	body string,
	metadata map[string][]string,
) string {
	s.store.EnsurePushCollection(collection)
	now := time.Now().UTC().Format(time.RFC3339)
	hash := store.HashContent(body)
	fresh := s.store.ActiveDocumentHash(collection, path) != hash

	if fresh {
		s.store.InsertContent(hash, body, now)
		s.store.InsertDocument(
			collection,
			path,
			store.ExtractTitle(body, path),
			hash,
			now,
		)
	}

	if len(metadata) > 0 {
		s.store.SetMetadata(collection, path, metadata)
	}

	if !fresh {
		return ""
	}

	chunks := chunk.Document(body, path, s.reranker)
	texts := make([]string, len(chunks))

	for i, c := range chunks {
		texts[i] = c.Text
	}

	embeddings, e := s.embedTexts(x, texts)

	if e != nil {
		return fmt.Sprintf(constant.EmbeddingPending, e)
	}

	for i, embedding := range embeddings {
		s.store.InsertEmbedding(hash, i, chunks[i].Position, embedding, now)
	}

	return ""
}
