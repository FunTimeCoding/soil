package channel

func (s *Server) WithSink(k Sink) *Server {
	s.sink = k

	return s
}
