package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/service"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Ansible,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d, constant.RunsTable)
	s := store.New(d, constant.RunsTable)
	s.AbandonRunning()
	p := reaper.New(r, y)
	result := &Daemon{
		database: d,
		reaper:   p,
		runner:   runner.New(o, s, l, r, p),
	}

	if o.Inventory != nil {
		c := o.Inventory()
		result.inventory = c
		result.inventoryGuard = peer.New(
			join.Colon(constant.InventoryPeer, c.Address()),
			c.Address(),
			c,
			l,
			r,
		)
		result.runner.WithInventory(guardedInventory(c, result.inventoryGuard))
	}

	result.service = service.New(s, result.runner)

	return result
}
