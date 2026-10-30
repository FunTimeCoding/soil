package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
)

func (s *Service) EmbedPending(x context.Context) error {
	for _, p := range s.store.PendingEmbeddings() {
		_, e := s.embedDocument(x, p.Hash, p.Path, p.Body)

		if e != nil && !validation.Is(e) {
			return e
		}
	}

	return nil
}
