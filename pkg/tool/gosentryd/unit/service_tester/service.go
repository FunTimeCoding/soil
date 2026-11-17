package service_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/service"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	g := peer.New(
		constant.SentryPeer,
		"sentry.test",
		o.Source,
		logger.New(context.Background()),
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(o.Source, g, "test")

	return o.service
}
