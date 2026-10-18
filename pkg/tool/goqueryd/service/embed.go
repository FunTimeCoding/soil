package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/embed"
)

func (s *Service) Embed(x context.Context) (*embed.Embed, error) {
	pending := s.store.PendingEmbeddings()
	r := embed.New(len(pending))

	for _, p := range pending {
		chunks, e := s.embedDocument(x, p)

		if e != nil {
			return r, e
		}

		r.Chunks += chunks
	}

	return r, nil
}
