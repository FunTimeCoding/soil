package offline_service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	gomemorydConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/workspace"
)

func New() (*workspace.Service, *commit.Service, func()) {
	l := logger.New(context.Background())
	r := memory.New()
	m := mock_client.New()
	g := peer.New(gomemorydConstant.Identity.Name(), m.Address(), m, l, r)
	w := workspace.New(index.DefaultDirectory())

	return w,
		commit.New(inventory.New(), w, m, g, r),
		func() {
			g.Close()
			m.Close()
		}
}
