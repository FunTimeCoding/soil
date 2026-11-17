package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/daemon"
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
		metric.New(),
	)
	o.t.Cleanup(o.daemon.Close)

	return o.daemon
}
