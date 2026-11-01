package socket

func (s *Socket) Close() {
	if w := s.client.WebSocket(); w != nil {
		w.Close()
	}
}
