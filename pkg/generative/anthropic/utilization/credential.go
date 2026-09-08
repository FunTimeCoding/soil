package utilization

import "time"

type Credential struct {
	AccessToken      string
	SubscriptionType string
	RateLimitTier    string
	ExpiresAt        time.Time
}
