package mock_system

func (s *System) Fail(
	command string,
	e error,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.failures[command] = e
}
