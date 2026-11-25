package keychain_store

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system/keychain"
	"golang.org/x/oauth2"
)

func (s *Store) SaveToken(t *oauth2.Token) error {
	if t.RefreshToken == "" {
		return validation.New("graph token carries no refresh token")
	}

	return keychain.Save(s.service, s.account, t.RefreshToken)
}
