package store

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/tool/gogated/model/login_session"
)

func (s *Store) CreateLoginSession(authorizeRequest string) (string, error) {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	identifier := hex.EncodeToString(b)
	e := s.mapper.Create(login_session.New(identifier, authorizeRequest)).Error

	return identifier, e
}
