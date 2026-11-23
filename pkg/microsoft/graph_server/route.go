package graph_server

func (s *Server) Route(
	path string,
	status int,
	body string,
) *Server {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.answers[path] = append(
		s.answers[path],
		Answer{Status: status, Body: body},
	)

	return s
}
