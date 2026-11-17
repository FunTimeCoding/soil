package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	goraidparsedConstant "github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	l := logger.New(o.t.Context())
	d := relational.Open(l, "", o.litePath)
	o.t.Cleanup(func() { relational.CloseMapper(d) })
	g := peer.New(
		goraidparsedConstant.Identity.Name(),
		"",
		o.Parser,
		l,
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(
		store.New(d),
		o.Parser,
		g,
		o.LogCache,
		o.Elite,
		o.Output,
		o.clock,
		l,
	)

	return o.service
}
