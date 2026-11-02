package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/worker"
)

func New(
	o *option.Sprout,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := lite.New(l, o.LitePath)
	migrate.AutoMigrate(d)
	n := o.Notifier()
	v := service.New(store.New(d, o.Clock), n)

	return &Daemon{
		database: d,
		service:  v,
		web:      web.New(v, n, r),
		watcher:  worker.New(v.Sync, o.SeedDirectory, l, r),
	}
}
