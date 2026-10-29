package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Netbox,
	l *logger.Logger,
	y prometheus.ObserverVec,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)

	return &Service{database: d, client: o.Source(y), store: store.New(d)}
}
