package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
	"time"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(
		t,
		service_tester.WithReporter(r),
		service_tester.WithDownMaster(),
		service_tester.WithKeySyncInterval(20*time.Millisecond),
	)
	workers := s.Daemon().Workers()
	workers[len(workers)-1].Close()

	return s.Served(), outage_tester.New(t).Guard(
		workers,
		constant.SaltPeer,
	).Unavailable
}
