package web

func (s *Server) multiple() bool {
	return len(s.service.Inventory().Instances) > 1
}
