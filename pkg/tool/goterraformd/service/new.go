package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Terraform,
	l *logger.Logger,
	r face.Reporter,
	y *prometheus.Registry,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d, constant.RunsTable)
	s := store.New(d, constant.RunsTable)
	p := reaper.New(r)

	return &Service{
		database: d,
		store:    s,
		reaper:   p,
		runner:   runner.New(o, s, l, r, p, y, o.Cluster()),
	}
}
