package supervisor

func (s *Supervisor) Restart(name string) error {
	s.commandMutex.Lock()
	defer s.commandMutex.Unlock()
	p, e := s.process(name)

	if e != nil {
		return e
	}

	if f := p.Stop(); f != nil {
		return f
	}

	s.spawn(p)

	return nil
}
