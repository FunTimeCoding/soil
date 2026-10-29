package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func New(
	o *option.Query,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := connection.New(l, o.LitePath)
	m := o.Embedder()

	return &Service{
		database: d,
		store:    store.New(d, o.EmbedModel),
		embedder: m,
		guard:    peer.New(constant.EmbedderPeer, o.EmbedTarget, m, l, r),
		reranker: o.Reranker(),
	}
}
