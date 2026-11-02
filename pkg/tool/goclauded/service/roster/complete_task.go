package roster

func (s *Service) CompleteTask(name string) (string, error) {
	topic, e := s.store.CompleteTask(name)

	if e != nil {
		return "", e
	}

	if topic != "" {
		s.notify()
	}

	return topic, nil
}
