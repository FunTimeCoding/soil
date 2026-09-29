package service

func (s *Service) resolveTag(name string) (int, error) {
	t, e := s.client.TagByName(name)

	if e != nil {
		return 0, e
	}

	return t.Identifier, nil
}
