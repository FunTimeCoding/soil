package stopper

func (s *Stopper) Close() {
	s.once.Do(func() { close(s.done) })
}
