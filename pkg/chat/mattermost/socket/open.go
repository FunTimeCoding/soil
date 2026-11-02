package socket

import "context"

func (s *Socket) Open(x context.Context) error {
	if s.opened || s.client.WebSocket() == nil {
		if e := s.client.RefreshSocket(x); e != nil {
			return e
		}
	}

	s.opened = true
	s.client.WebSocket().Listen()

	return nil
}
