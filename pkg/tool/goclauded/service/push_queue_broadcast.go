package service

func (s *Service) PushQueueBroadcast(
	kind string,
	body string,
) error {
	sessions, e := s.store.ListSessions()

	if e != nil {
		return e
	}

	if e := s.store.PushQueueBroadcast(sessions, kind, body); e != nil {
		return e
	}

	s.notify()

	return nil
}
