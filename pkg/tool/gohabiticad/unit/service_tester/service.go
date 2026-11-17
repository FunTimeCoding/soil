package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/service"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	g := peer.New(
		constant.HabiticaPeer,
		"habitica.test",
		o.Source,
		logger.New(context.Background()),
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(o.Source, g)

	return o.service
}
