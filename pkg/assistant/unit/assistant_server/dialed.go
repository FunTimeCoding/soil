package assistant_server

func (s *Server) Dialed() <-chan struct{} {
	return s.dialed
}
