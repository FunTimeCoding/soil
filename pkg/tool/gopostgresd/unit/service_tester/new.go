package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/postgres/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/dump/mock_dump"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/peered"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/backup"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/catalog"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/query"
	backupStore "github.com/funtimecoding/soil/pkg/tool/gopostgresd/store/backup"
	"testing"
	"time"
)

func New(
	t *testing.T,
	names ...string,
) (*Tester, map[string]*mock_client.Client) {
	t.Helper()
	clients := make(map[string]*mock_client.Client)
	var instances []instance.Instance

	for _, name := range names {
		clients[name] = mock_client.New()
		instances = append(
			instances,
			instance.Instance{
				Name:     name,
				Host:     join.Empty(name, ".example"),
				Database: name,
			},
		)
	}

	l := logger.New(t.Context())
	r := memory.New()
	i := inventory.New[instance.Instance](instances...)
	p := pool.New(
		i,
		func(i *instance.Instance) (*peered.Peered, error) {
			c := clients[i.Name]

			return peered.New(
				c,
				peer.New(
					join.Colon(constant.InstancePeer, i.Name),
					i.Host,
					c,
					l,
					r,
				),
			), nil
		},
	)
	t.Cleanup(p.Close)
	result := &Tester{
		Catalog:   catalog.New(p),
		Query:     query.New(p),
		Dump:      mock_dump.New(),
		Directory: t.TempDir(),
	}
	result.Backup = backup.New(
		i,
		p,
		backupStore.New(result.Directory),
		result.Dump,
		func() time.Time { return Now() },
	)

	return result, clients
}
