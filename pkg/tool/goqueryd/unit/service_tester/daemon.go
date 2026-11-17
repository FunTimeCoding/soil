package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/daemon"
	"github.com/prometheus/client_golang/prometheus"
)

func (o *Tester) Daemon() *daemon.Daemon {
	o.t.Helper()

	if o.daemon != nil {
		return o.daemon
	}

	o.daemon = daemon.New(
		o.daemonOption(),
		logger.New(o.t.Context()),
		o.Reporter,
		prometheus.NewRegistry(),
	)
	o.t.Cleanup(o.daemon.Close)

	return o.daemon
}
