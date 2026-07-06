package ticker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"time"
)

func New(
	i time.Duration,
	f func(),
	r *recovery.Recovery,
) *Ticker {
	return &Ticker{interval: i, function: f, recovery: r}
}
