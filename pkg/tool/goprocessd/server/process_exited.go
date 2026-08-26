package server

func (s *Server) processExited(supervisorStopped bool) {
	s.countMutex.Lock()
	s.running--
	remaining := s.running
	s.countMutex.Unlock()

	if supervisorStopped || remaining > 0 {
		return
	}

	select {
	case s.allDone <- struct{}{}:
	default:
	}
}
