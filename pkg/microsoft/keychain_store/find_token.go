package keychain_store

import (
	"github.com/funtimecoding/soil/pkg/system/keychain"
	"golang.org/x/oauth2"
)

func (s *Store) FindToken() (*oauth2.Token, bool, error) {
	refresh, found, e := keychain.Find(s.service, s.account)

	if e != nil || !found {
		return nil, false, e
	}

	return &oauth2.Token{RefreshToken: refresh}, true, nil
}
