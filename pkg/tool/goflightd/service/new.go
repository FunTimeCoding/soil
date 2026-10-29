package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/option"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
)

func New(
	o *option.Flight,
	l *logger.Logger,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)

	return &Service{
		database:  d,
		store:     store.New(d),
		logger:    l,
		predicate: o.Predicate,
		retention: o.Retention,
		clock:     o.Clock,
		command:   o.Command,
		stream:    o.Stream,
		last:      make(map[string]string),
	}
}
