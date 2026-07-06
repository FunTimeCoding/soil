package ticker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"time"
)

type Ticker struct {
	interval time.Duration
	function func()
	recovery *recovery.Recovery
	ticker   *time.Ticker
	done     chan struct{}
}
