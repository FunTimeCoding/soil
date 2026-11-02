package session

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func (s *Session) DeleteHash() string {
	hash := sha256.Sum256(
		[]byte(
			fmt.Sprintf(
				"%s|%d|%s",
				s.Identifier,
				s.TurnCount,
				s.LastActiveAt.UTC().Format(time.RFC3339Nano),
			),
		),
	)

	return fmt.Sprintf("%x", hash[:4])
}
