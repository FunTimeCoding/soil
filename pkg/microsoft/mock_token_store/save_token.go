package mock_token_store

import "golang.org/x/oauth2"

func (s *Store) SaveToken(t *oauth2.Token) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.Failure != nil {
		return s.Failure
	}

	if s.SaveFailure != nil {
		return s.SaveFailure
	}

	saved := *t
	s.token = &saved
	s.saved = append(s.saved, &saved)

	return nil
}
