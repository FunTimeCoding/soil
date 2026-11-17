package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/service"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service == nil {
		g := peer.New(
			constant.OpenbaoPeer,
			o.Source.Address(),
			o.Source,
			logger.New(o.t.Context()),
			o.Reporter,
		)
		o.t.Cleanup(g.Close)
		o.service = service.New(o.Source, g)
	}

	return o.service
}
