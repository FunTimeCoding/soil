package session

import "context"

func (s *Session) Renew() error {
	return s.Login(context.Background())
}
