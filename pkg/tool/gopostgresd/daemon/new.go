package daemon

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/option"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/backup"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/catalog"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service/query"
	backupStore "github.com/funtimecoding/soil/pkg/tool/gopostgresd/store/backup"
)

func New(
	o *option.Postgres,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	p := pool.New(o.Inventory, peeredBuilder(o, l, r))

	for _, i := range o.Inventory.Instances {
		_, e := p.Client(i.Name)
		errors.PanicOnError(e)
	}

	return &Daemon{
		inventory: o.Inventory,
		pool:      p,
		selection: selection.New(),
		catalog:   catalog.New(p),
		query:     query.New(p),
		backup: backup.New(
			o.Inventory,
			p,
			backupStore.New(o.BackupDirectory),
			o.Dump,
			o.Clock,
		),
	}
}
