package mock_token_store

import "golang.org/x/oauth2"

func (s *Store) Saved() []*oauth2.Token {
	s.lock.Lock()
	defer s.lock.Unlock()

	return append([]*oauth2.Token{}, s.saved...)
}
