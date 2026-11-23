package keychain_store

import (
	"encoding/base64"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/system/keychain"
	"golang.org/x/oauth2"
)

func (s *Store) SaveToken(t *oauth2.Token) error {
	b, e := json.Marshal(t)

	if e != nil {
		return e
	}

	return keychain.Save(
		s.service,
		s.account,
		base64.StdEncoding.EncodeToString(b),
	)
}
