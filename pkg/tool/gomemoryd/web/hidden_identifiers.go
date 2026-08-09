package web

func (s *Server) hiddenIdentifiers() map[int64]bool {
	hidden, e := s.service.HiddenIdentifiers()

	if e != nil {
		return map[int64]bool{}
	}

	return hidden
}
