package sweeper

func (s *Sweeper) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.running {
		return
	}

	s.running = false
	close(s.done)
}
