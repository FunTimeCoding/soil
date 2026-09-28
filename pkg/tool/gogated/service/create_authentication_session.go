package service

import "time"

func (s *Service) CreateAuthenticationSession(
	userIdentifier string,
	userAgent string,
	address string,
	authenticatedAt time.Time,
) (string, error) {
	return s.store.CreateAuthenticationSession(
		userIdentifier,
		userAgent,
		address,
		authenticatedAt,
	)
}
