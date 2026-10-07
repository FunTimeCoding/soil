package credential

import "time"

func New(
	accessToken string,
	subscriptionType string,
	rateLimitTier string,
	expiresAt time.Time,
) *Credential {
	return &Credential{
		AccessToken:      accessToken,
		SubscriptionType: subscriptionType,
		RateLimitTier:    rateLimitTier,
		ExpiresAt:        expiresAt,
	}
}
