package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/base"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(t, base.WithReporter(r))
	s.UnreachableInventory()
	workers := s.Daemon().Workers()
	workers[len(workers)-1].Close()

	return s.Served(), outage_tester.Guard(
		t,
		workers,
		join.Colon(
			constant.InventoryPeer,
			locator.New("localhost").Port(1).Insecure().Address(),
		),
	).Unavailable
}
