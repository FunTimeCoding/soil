package base

func (s *Server) CloseDaemon() {
	s.daemon.Close()
}
