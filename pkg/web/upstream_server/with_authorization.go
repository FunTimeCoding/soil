package upstream_server

func WithAuthorization(accepted string) Option {
	return func(s *Server) {
		s.authorization = accepted
	}
}
