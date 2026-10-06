package unit

func (s *saltServer) expire() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.valid = "expired"
}
