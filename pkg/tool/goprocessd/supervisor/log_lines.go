package supervisor

func (s *Supervisor) LogLines(
	name string,
	all bool,
) ([]string, int, error) {
	p, e := s.process(name)

	if e != nil {
		return nil, 0, e
	}

	if all {
		return p.Log(), 0, nil
	}

	current, older := p.CurrentLog()

	return current, older, nil
}
