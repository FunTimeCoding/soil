package service

func (s *Service) Release(
	item string,
	owner string,
) error {
	if e := validateClaim(item, owner); e != nil {
		return e
	}

	return s.store.Release(item, owner)
}
