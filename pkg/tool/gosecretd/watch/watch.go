package watch

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/ticker"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
	"github.com/prometheus/client_golang/prometheus"
)

type Watch struct {
	source    face.SecretSource
	sealed    prometheus.Gauge
	reachable prometheus.Gauge
	recovery  *recovery.Recovery
	ticker    *ticker.Ticker
}
