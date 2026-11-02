package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/worker"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Query,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Daemon {
	d := connection.New(l, o.LitePath)
	m := o.Embedder()
	result := &Daemon{
		database: d,
		embedder: m,
		guard:    peer.New(constant.EmbedderPeer, o.EmbedTarget, m, l, r),
		reranker: o.Reranker(),
		counter:  o.EmbedCounter(),
	}
	result.service = service.New(
		o.EmbedModel,
		store.New(d, o.EmbedModel),
		m,
		result.guard,
		result.reranker,
		result.counter,
	)
	result.web = web.New(result.service, r)
	result.clock = worker.New(
		result.service.Maintain,
		constant.MaintainInterval,
		l,
		y,
		r,
	)

	return result
}
