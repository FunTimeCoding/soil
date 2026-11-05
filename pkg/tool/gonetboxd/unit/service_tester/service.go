package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"path/filepath"
	"testing"
)

func Service(
	t *testing.T,
	c face.NetBoxClient,
) *service.Service {
	t.Helper()
	l := logger.New(t.Context())
	d := relational.Open(l, "", filepath.Join(t.TempDir(), "gonetboxd.db"))
	migrate.AutoMigrate(d)
	t.Cleanup(func() { relational.CloseMapper(d) })
	g := peer.New(constant.NetBoxPeer, c.Address(), c, l, memory.New())
	t.Cleanup(g.Close)

	return service.New(store.New(d), c, g)
}
