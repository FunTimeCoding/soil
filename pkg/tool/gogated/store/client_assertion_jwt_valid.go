package store

import "context"

func (s *Store) ClientAssertionJWTValid(
	_ context.Context,
	_ string,
) error {
	return nil
}
