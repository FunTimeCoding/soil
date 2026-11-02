package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/worker"
	"testing"
)

func NewWatcher(
	t *testing.T,
	v *service.Service,
	r face.Reporter,
) *worker.Worker {
	t.Helper()

	return worker.New(
		v.HandleEvent,
		v.WatchedPaths(),
		logger.New(t.Context()),
		r,
		v.SyncLogCache,
		v.EnrichExisting,
		v.Cleanup,
	)
}
