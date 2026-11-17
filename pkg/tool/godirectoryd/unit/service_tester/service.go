package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	g := peer.New(
		constant.DirectoryPeer,
		"directory.test",
		o.Source,
		logger.New(o.t.Context()),
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(o.Source, g)

	return o.service
}
