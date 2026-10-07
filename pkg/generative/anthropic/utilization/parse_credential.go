package utilization

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"
	"time"
)

func ParseCredential(raw string) *credential.Credential {
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

	return credential.New(
		stored.Claude.AccessToken,
		stored.Claude.SubscriptionType,
		stored.Claude.RateLimitTier,
		time.UnixMilli(stored.Claude.ExpiresAt),
	)
}
