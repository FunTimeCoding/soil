package utilization

type claudeCredential struct {
	AccessToken      string `json:"accessToken"`
	SubscriptionType string `json:"subscriptionType"`
	RateLimitTier    string `json:"rateLimitTier"`
	ExpiresAt        int64  `json:"expiresAt"`
}
