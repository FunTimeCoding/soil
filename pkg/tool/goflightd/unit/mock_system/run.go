package mock_system

func (s *System) Run(v ...string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if e := s.failures[v[0]]; e != nil {
		return "", e
	}

	return s.outputs[v[0]], nil
}
