package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/service"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service == nil {
		g := peer.New(
			constant.OpnsensePeer,
			o.source.Address(),
			o.source,
			logger.New(o.t.Context()),
			o.Reporter,
		)
		o.t.Cleanup(g.Close)
		o.service = service.New(o.source, g)
	}

	return o.service
}
