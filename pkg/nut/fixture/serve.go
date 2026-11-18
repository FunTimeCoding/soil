package fixture

func (s *Server) serve() {
	for {
		n, e := s.listener.Accept()

		if e != nil {
			return
		}

		go s.answer(n)
	}
}
