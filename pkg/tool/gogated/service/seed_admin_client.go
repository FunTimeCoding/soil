package service

func (s *Service) SeedAdminClient(
	identifier string,
	secret string,
	issuer string,
) error {
	return s.store.SeedAdminClient(identifier, secret, issuer)
}
