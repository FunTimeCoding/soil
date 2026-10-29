package service

import "context"

func (s *Service) EmbedPending(x context.Context) error {
	for _, p := range s.store.PendingEmbeddings() {
		if _, e := s.embedDocument(x, p); e != nil {
			return e
		}
	}

	return nil
}
