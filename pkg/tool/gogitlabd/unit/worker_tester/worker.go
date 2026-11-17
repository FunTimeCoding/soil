package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func (o *Tester) Worker() *worker.Worker {
	o.t.Helper()

	if o.worker != nil {
		return o.worker
	}

	s := o.Service()
	o.worker = worker.New(
		s.Refresh,
		s.Entries,
		time.Hour,
		prometheus.NewRegistry(),
		logger.New(o.t.Context()),
		o.Reporter,
	)
	o.t.Cleanup(o.worker.Close)

	return o.worker
}
