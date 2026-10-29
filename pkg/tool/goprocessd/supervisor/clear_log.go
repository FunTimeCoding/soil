package supervisor

func (s *Supervisor) ClearLog(name string) error {
	p, e := s.process(name)

	if e != nil {
		return e
	}

	p.ClearLog()

	return nil
}
