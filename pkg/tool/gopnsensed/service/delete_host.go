package service

func (s *Service) DeleteHost(
	identifier string,
	apply bool,
) error {
	return s.attempt(
		func() error {
			if e := s.client.DeleteHost(identifier); e != nil {
				return e
			}

			return s.reconfigureWhen(apply)
		},
	)
}
