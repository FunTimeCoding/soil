package bidi_server

func (s *Server) Stop() {
	s.server.CloseClientConnections()
	s.server.Close()
}
