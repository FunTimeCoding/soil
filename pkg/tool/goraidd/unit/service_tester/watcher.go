package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/worker"
)

func (o *Tester) Watcher(v *service.Service) *worker.Worker {
	o.t.Helper()

	return worker.New(
		v.HandleEvent,
		v.WatchedPaths(),
		logger.New(o.t.Context()),
		o.Reporter,
		v.SyncLogCache,
		v.EnrichExisting,
		v.Cleanup,
	)
}
