package authentication_session

import "time"

func New(
	identifier string,
	userIdentifier string,
	userAgent string,
	address string,
	authenticatedAt time.Time,
) *AuthenticationSession {
	return &AuthenticationSession{
		Identifier:      identifier,
		UserIdentifier:  userIdentifier,
		UserAgent:       userAgent,
		Address:         address,
		AuthenticatedAt: authenticatedAt,
		LastUsedAt:      authenticatedAt,
	}
}
