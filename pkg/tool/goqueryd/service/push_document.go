package service

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"time"
)

func (s *Service) PushDocument(
	collection string,
	path string,
	body string,
	metadata map[string]string,
) error {
	s.store.EnsurePushCollection(collection)
	now := time.Now().UTC().Format(time.RFC3339)
	hash := store.HashContent(body)
	title := store.ExtractTitle(body, path)
	s.store.InsertContent(hash, body, now)
	s.store.InsertDocument(collection, path, title, hash, now)
	chunks := chunk.Document(body, path)
	texts := make([]string, len(chunks))

	for i, c := range chunks {
		texts[i] = c.Text
	}

	embeddings, e := s.ollama.Embed(constant.EmbedModel, texts)

	if e != nil {
		return e
	}

	for i, embedding := range embeddings {
		s.store.InsertEmbedding(
			hash,
			i,
			chunks[i].Position,
			embedding,
			now,
		)
	}

	if len(metadata) > 0 {
		s.store.SetMetadata(collection, path, metadata)
	}

	return nil
}
