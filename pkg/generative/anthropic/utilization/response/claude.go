package response

type Claude struct {
	AccessToken      string `json:"accessToken"`
	SubscriptionType string `json:"subscriptionType"`
	RateLimitTier    string `json:"rateLimitTier"`
	ExpiresAt        int64  `json:"expiresAt"`
}
