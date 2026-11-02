package server_tester

func (s *Server) serve() {
	defer s.group.Done()

	for {
		c, e := s.listener.Accept()

		if e != nil {
			return
		}

		s.mutex.Lock()
		s.connections = append(s.connections, c)
		s.mutex.Unlock()
		s.group.Add(1)
		go s.serveConnection(c)
	}
}
