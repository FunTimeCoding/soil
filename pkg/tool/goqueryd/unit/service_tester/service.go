package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
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
		constant.EmbedderPeer,
		constant.FixtureEmbedTarget,
		o.embedder,
		l,
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(
		o.EmbedModel,
		store.New(d, o.EmbedModel),
		o.embedder,
		g,
		o.reranker,
		o.counter,
	)

	return o.service
}
