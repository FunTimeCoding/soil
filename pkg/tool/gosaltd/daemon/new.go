package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/service"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Salt,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d, constant.RunsTable)
	s := store.New(d, constant.RunsTable)
	s.AbandonRunning()
	p := reaper.New(r, y)
	c := o.Master()
	result := &Daemon{
		database: d,
		reaper:   p,
		salt:     c,
		runner:   runner.New(o, s, l, r, p, c),
	}
	result.service = service.New(s, result.runner, c)

	return result
}
