package service

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func (s *Service) DeleteHash(identifier string) string {
	r, found, e := s.store.FindSession(identifier)

	if e != nil || !found {
		return ""
	}

	hash := sha256.Sum256(
		[]byte(
			fmt.Sprintf(
				"%s|%d|%s",
				r.Identifier,
				r.TurnCount,
				r.LastActiveAt.UTC().Format(time.RFC3339Nano),
			),
		),
	)

	return fmt.Sprintf("%x", hash[:4])
}
