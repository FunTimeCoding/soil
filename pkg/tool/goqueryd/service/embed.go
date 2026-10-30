package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/embed"
)

func (s *Service) Embed(x context.Context) (*embed.Embed, error) {
	pending := s.store.PendingEmbeddings()
	r := embed.New(len(pending))
	var refused error

	for _, p := range pending {
		chunks, e := s.embedDocument(x, p.Hash, p.Path, p.Body)
		r.Chunks += chunks

		if e == nil {
			continue
		}

		if !validation.Is(e) {
			return r, e
		}

		if refused == nil {
			refused = e
		}
	}

	return r, refused
}
