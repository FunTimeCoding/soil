package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	memoryConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"testing"
)

func fromParts(
	t *testing.T,
	i *inventory.Inventory,
	m *mock_client.Client,
) *service.Service {
	t.Helper()
	l := logger.New(t.Context())
	r := memory.New()
	g := peer.New(memoryConstant.Identity.Name(), m.Address(), m, l, r)
	t.Cleanup(g.Close)

	return service.New(i, index.DefaultDirectory(), m, g, r)
}
