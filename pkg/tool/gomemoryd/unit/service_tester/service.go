package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	goquerydConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	l := logger.New(o.t.Context())
	d := connection.New(l, o.Path)
	o.t.Cleanup(func() { errors.PanicClose(d) })
	g := peer.New(
		goquerydConstant.Identity.Name(),
		o.Indexer.Address(),
		o.Indexer,
		l,
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(
		store.New(d),
		o.Indexer,
		g,
		o.hiddenTag,
		o.referenceRoot,
		o.clock,
		o.Reporter,
	)

	return o.service
}
