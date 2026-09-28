package store

import (
	"context"
	"time"
)

func (s *Store) SetClientAssertionJWT(
	_ context.Context,
	_ string,
	_ time.Time,
) error {
	return nil
}
