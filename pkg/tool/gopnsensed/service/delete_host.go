package service

func (s *Service) DeleteHost(
	identifier string,
	apply bool,
) error {
	if e := s.source.DeleteHost(identifier); e != nil {
		return e
	}

	return s.reconfigureWhen(apply)
}
