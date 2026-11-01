package socket

import "context"

func (s *Socket) Open(context.Context) error {
	if s.opened || s.client.WebSocket() == nil {
		if e := s.client.RefreshSocket(); e != nil {
			return e
		}
	}

	s.opened = true
	s.client.WebSocket().Listen()

	return nil
}
