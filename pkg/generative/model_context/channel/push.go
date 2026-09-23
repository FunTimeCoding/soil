package channel

func (s *Server) Push(
	content string,
	meta map[string]string,
) {
	s.sink.Push(content, meta)
}
