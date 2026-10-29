package service

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
)

func New(
	o *option.Sprout,
	l *logger.Logger,
) *Service {
	d := lite.New(l, o.LitePath)
	migrate.AutoMigrate(d)

	return &Service{
		database: d,
		store:    store.New(d, o.Clock),
		notifier: o.Notifier(),
	}
}
