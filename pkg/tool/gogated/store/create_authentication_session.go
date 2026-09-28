package store

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/authentication_session"
	"time"
)

func (s *Store) CreateAuthenticationSession(
	userIdentifier string,
	userAgent string,
	address string,
	authenticatedAt time.Time,
) (string, error) {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	identifier := hex.EncodeToString(b)
	e := s.mapper.Create(
		authentication_session.New(
			identifier,
			userIdentifier,
			userAgent,
			address,
			authenticatedAt,
		),
	).Error

	return identifier, e
}
