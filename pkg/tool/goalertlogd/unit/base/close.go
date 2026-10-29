package base

func (s *Server) Close() {
	s.Server.Close()
	s.Worker.Close()
	s.MockClient.Close()
}
