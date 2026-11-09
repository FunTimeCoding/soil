package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
	goraidparsedConstant "github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
	"testing"
)

func NewService(
	t *testing.T,
	o *option.Raid,
	r face.Reporter,
) *service.Service {
	t.Helper()
	l := logger.New(t.Context())
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(d) })
	p := o.Parser()
	g := peer.New(goraidparsedConstant.Identity.Name(), "", p, l, r)
	t.Cleanup(g.Close)

	return service.New(
		store.New(d),
		p,
		g,
		o.LogCachePath,
		o.ElitePath,
		o.OutputPath,
		o.Clock,
		l,
	)
}
