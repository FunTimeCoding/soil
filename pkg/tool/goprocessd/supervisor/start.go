package supervisor

func (s *Supervisor) Start(name string) error {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()
	p, e := s.process(name)

	if e != nil {
		return e
	}

	s.spawn(p)

	return nil
}
