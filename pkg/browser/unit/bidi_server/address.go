package bidi_server

func (s *Server) Address() string {
	return s.server.Listener.Addr().String()
}
