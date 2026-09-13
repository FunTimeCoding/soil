package store

func (s *Store) notify() {
	if s.notifier == nil {
		return
	}

	s.notifier.Notify()
}
