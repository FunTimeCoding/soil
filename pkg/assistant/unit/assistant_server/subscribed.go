package assistant_server

func (s *Server) Subscribed() <-chan struct{} {
	return s.subscribed
}
