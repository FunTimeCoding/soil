package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"time"
)

func (o *Tester) Worker() *worker.Worker {
	o.t.Helper()

	if o.worker != nil {
		return o.worker
	}

	o.worker = worker.New(
		o.Service().Refresh,
		time.Minute,
		o.Registry,
		logger.New(o.t.Context()),
		o.Reporter,
	)
	o.t.Cleanup(o.worker.Close)

	return o.worker
}
