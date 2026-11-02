package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func New(
	t *testing.T,
	o *option.Query,
	r face.Reporter,
) *service.Service {
	t.Helper()
	l := logger.New(t.Context())
	d := connection.New(l, o.LitePath)
	t.Cleanup(func() { errors.PanicClose(d) })
	m := o.Embedder()
	t.Cleanup(m.Close)
	g := peer.New(constant.EmbedderPeer, o.EmbedTarget, m, l, r)
	t.Cleanup(g.Close)
	a := o.Reranker()
	t.Cleanup(func() { errors.LogClose(a) })
	c := o.EmbedCounter()
	t.Cleanup(func() { errors.LogClose(c) })

	return service.New(o.EmbedModel, store.New(d, o.EmbedModel), m, g, a, c)
}
