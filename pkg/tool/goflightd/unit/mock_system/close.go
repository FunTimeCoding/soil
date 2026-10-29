package mock_system

func (s *System) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.stream != nil {
		s.stream.Close()
	}
}
