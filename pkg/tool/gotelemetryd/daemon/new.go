package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/web"
)

func New(
	o *option.Telemetry,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	v := service.New(store.New(d), o.Clock, r)

	return &Daemon{database: d, service: v, web: web.New(v)}
}
