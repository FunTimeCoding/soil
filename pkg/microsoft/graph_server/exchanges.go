package graph_server

func (s *Server) Exchanges(path string) []Exchange {
	s.lock.Lock()
	defer s.lock.Unlock()
	var result []Exchange

	for _, x := range s.exchanges {
		if x.Path == path {
			result = append(result, x)
		}
	}

	return result
}
