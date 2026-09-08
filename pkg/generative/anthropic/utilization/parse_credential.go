package utilization

import (
	"encoding/json"
	"time"
)

func ParseCredential(raw string) *Credential {
	if raw == "" {
		return nil
	}

	var stored storedCredential

	if e := json.Unmarshal([]byte(raw), &stored); e != nil {
		return nil
	}

	if stored.Claude.AccessToken == "" {
		return nil
	}

	return &Credential{
		AccessToken:      stored.Claude.AccessToken,
		SubscriptionType: stored.Claude.SubscriptionType,
		RateLimitTier:    stored.Claude.RateLimitTier,
		ExpiresAt:        time.UnixMilli(stored.Claude.ExpiresAt),
	}
}
