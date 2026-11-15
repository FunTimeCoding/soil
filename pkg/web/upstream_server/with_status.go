package upstream_server

func WithStatus(codes ...int) Option {
	return func(s *Server) {
		s.statuses = codes
	}
}
