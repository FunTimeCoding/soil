package daemon

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/option"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/service"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
)

func New(
	o *option.Monitor,
	l *logger.Logger,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	n := notifier.New()

	return &Daemon{database: d, service: service.New(store.New(d, n), n)}
}
