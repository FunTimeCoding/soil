package model_context_server

func WithSignOn() Option {
	return func(s *Server) {
		s.signOn = true
	}
}
