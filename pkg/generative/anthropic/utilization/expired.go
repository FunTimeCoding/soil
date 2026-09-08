package utilization

import "time"

func (c *Credential) Expired(now time.Time) bool {
	return !c.ExpiresAt.After(now)
}
