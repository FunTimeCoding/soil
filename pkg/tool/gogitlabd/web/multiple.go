package web

func (s *Server) multiple() bool {
	return len(s.inventory.Instances) > 1
}
