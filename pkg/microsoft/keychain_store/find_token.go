package keychain_store

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/keychain"
	"golang.org/x/oauth2"
)

func (s *Store) FindToken() (*oauth2.Token, bool, error) {
	encoded, found, e := keychain.Find(s.service, s.account)

	if e != nil || !found {
		return nil, false, e
	}

	b, e := base64.StdEncoding.DecodeString(encoded)

	if e != nil {
		return nil, false, fmt.Errorf("decode keychain token: %w", e)
	}

	result := &oauth2.Token{}

	if f := json.Unmarshal(b, result); f != nil {
		return nil, false, fmt.Errorf("read keychain token: %w", f)
	}

	return result, true, nil
}
