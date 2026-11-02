package daemon

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web"
)

func New(
	o *option.Log,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	n := notifier.New()
	v := service.New(store.New(d), n)

	return &Daemon{database: d, service: v, web: web.New(v, n, r)}
}
