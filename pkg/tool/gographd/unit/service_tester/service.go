package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gographd/service"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	g := peer.New(
		constant.GraphPeer,
		microsoftConstant.GraphHost,
		o.Source,
		logger.New(context.Background()),
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(o.Source, g)

	return o.service
}
