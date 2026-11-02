package service

func (s *Service) Unsubscribe(
	callsign string,
	root string,
) (bool, error) {
	if e := s.requireSubscriptions(); e != nil {
		return false, e
	}

	count, e := s.store.Delete(callsign, root)

	if e != nil {
		return false, e
	}

	if count == 0 {
		return false, nil
	}

	return true, s.forgetIfUnwatched(root)
}
