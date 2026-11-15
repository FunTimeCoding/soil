package upstream_server

func WithTLS() Option {
	return func(s *Server) {
		s.secure = true
	}
}
