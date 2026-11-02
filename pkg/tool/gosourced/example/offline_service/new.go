package offline_service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	memoryConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
)

func New() (*service.Service, func()) {
	l := logger.New(context.Background())
	r := memory.New()
	m := mock_client.New()
	g := peer.New(memoryConstant.Identity.Name(), m.Address(), m, l, r)

	return service.New(inventory.New(), index.DefaultDirectory(), m, g, r),
		func() {
			g.Close()
			m.Close()
		}
}
