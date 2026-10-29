package supervisor

func (s *Supervisor) Stop(name string) error {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()
	p, e := s.process(name)

	if e != nil {
		return e
	}

	return p.Stop()
}
