package mock_system

func (s *System) Answer(
	command string,
	output string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.outputs[command] = output
}
