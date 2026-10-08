package unit

func (s *SaltServer) expire() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.valid = "expired"
}
