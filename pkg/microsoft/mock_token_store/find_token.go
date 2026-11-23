package mock_token_store

import "golang.org/x/oauth2"

func (s *Store) FindToken() (*oauth2.Token, bool, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.Failure != nil {
		return nil, false, s.Failure
	}

	if s.token == nil {
		return nil, false, nil
	}

	result := *s.token

	return &result, true, nil
}
