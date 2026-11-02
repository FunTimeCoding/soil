package service

import "context"

func (s *Service) Maintain(x context.Context) (int64, error) {
	if _, e := s.IndexCollections(""); e != nil {
		return 0, e
	}

	result := s.Prune()

	return result, s.EmbedPending(x)
}
