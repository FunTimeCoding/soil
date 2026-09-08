package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	"time"
)

func Utilization() {
	if !utilization.Supported() {
		console.Line("credential storage is not supported on this platform")

		return
	}

	c := utilization.ReadCredential()

	if c == nil {
		console.Line("no credential found")

		return
	}

	console.Format(
		"Subscription: %s  Tier: %s  Expires: %s\n",
		c.SubscriptionType,
		c.RateLimitTier,
		c.ExpiresAt.Local().Format(time.RFC3339),
	)

	if c.Expired(time.Now()) {
		console.Line("credential expired")

		return
	}

	r := utilization.Read(c.AccessToken)

	if r == nil {
		console.Line("no utilization data")

		return
	}

	console.Format(
		"Session %d%%  resets %s\n",
		r.SessionPercent,
		r.SessionReset.Local().Format(time.RFC3339),
	)
	console.Format(
		"Weekly  %d%%  resets %s\n",
		r.WeeklyPercent,
		r.WeeklyReset.Local().Format(time.RFC3339),
	)

	if !r.FableSeen {
		console.Line("Fable   not reported")

		return
	}

	if r.FableReset.IsZero() {
		console.Format("Fable   %d%%  never used\n", r.FablePercent)

		return
	}

	console.Format(
		"Fable   %d%%  resets %s\n",
		r.FablePercent,
		r.FableReset.Local().Format(time.RFC3339),
	)
}
