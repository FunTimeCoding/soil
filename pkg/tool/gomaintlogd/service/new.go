package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
)

func New(
	o *option.Log,
	l *logger.Logger,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)

	return &Service{
		database: d,
		store:    store.New(d),
		notifier: notifier.New(),
	}
}
