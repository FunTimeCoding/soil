package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/option"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
)

func New(
	o *option.Monitor,
	l *logger.Logger,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	n := notifier.New()

	return &Service{database: d, notifier: n, store: store.New(d, n)}
}
