package channel

func (s *Server) Resolve(callsign string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.callsign = callsign
}
