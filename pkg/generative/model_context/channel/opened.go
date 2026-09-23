package channel

func (s *Server) Opened() bool {
	select {
	case <-s.open:
		return true
	default:
		return false
	}
}
