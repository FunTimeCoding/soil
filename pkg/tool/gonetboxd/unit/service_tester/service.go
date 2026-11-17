package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

func (o *Tester) Service() *service.Service {
	o.t.Helper()

	if o.service != nil {
		return o.service
	}

	l := logger.New(o.t.Context())
	d := relational.Open(l, "", o.litePath)
	migrate.AutoMigrate(d)
	o.t.Cleanup(func() { relational.CloseMapper(d) })
	g := peer.New(
		constant.NetBoxPeer,
		o.client.Address(),
		o.client,
		l,
		o.Reporter,
	)
	o.t.Cleanup(g.Close)
	o.service = service.New(store.New(d), o.client, g)

	return o.service
}
