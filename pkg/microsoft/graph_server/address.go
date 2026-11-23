package graph_server

func (s *Server) Address() string {
	return s.server.URL
}
