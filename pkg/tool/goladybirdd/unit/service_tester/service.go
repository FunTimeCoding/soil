package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/navigation"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/service"
	"time"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service == nil {
		g := peer.New(
			constant.LadybirdPeer,
			o.Client.Address(),
			o.Client,
			logger.New(o.t.Context()),
			o.Reporter,
		)
		o.t.Cleanup(g.Close)
		w := navigation.New(50*time.Millisecond, 300*time.Millisecond)
		o.Client.Listen(w.Observe)
		o.service = service.New(o.Client, g, w, o.t.TempDir())
	}

	return o.service
}
