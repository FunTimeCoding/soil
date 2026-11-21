package score

func (s *Service) load() error {
	s.mutex.RLock()
	loaded := s.loaded
	s.mutex.RUnlock()

	if loaded {
		return nil
	}

	rows, e := s.store.Scores()

	if e != nil {
		return e
	}

	state, f := s.store.MirrorState()

	if f != nil {
		return f
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i := range rows {
		s.scores[rows[i].Identifier] = &rows[i]
	}

	s.state = state
	s.loaded = true
	s.generation++

	return nil
}
