package upstream_server

func WithAnswer(a string) Option {
	return func(s *Server) {
		s.answer = a
	}
}
