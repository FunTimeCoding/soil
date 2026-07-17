package service

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/embed"
	"time"
)

func (s *Service) Embed() (*embed.Embed, error) {
	pending := s.store.PendingEmbeddings()
	r := embed.New(len(pending))
	now := time.Now().UTC().Format(time.RFC3339)

	for _, p := range pending {
		chunks := chunk.Document(p.Body, p.Path)
		texts := make([]string, len(chunks))

		for i, c := range chunks {
			texts[i] = c.Text
		}

		embeddings, e := s.ollama.Embed(constant.EmbedModel, texts)

		if e != nil {
			return r, e
		}

		for i, embedding := range embeddings {
			s.store.InsertEmbedding(
				p.Hash,
				i,
				chunks[i].Position,
				embedding,
				now,
			)
		}

		r.Chunks += len(chunks)
	}

	return r, nil
}
