package utilization

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/response"
	"time"
)

func ParseCredential(raw string) *credential.Credential {
	if raw == "" {
		return nil
	}

	var stored response.Stored

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
