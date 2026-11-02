package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/migrate"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
)

func New(
	o *option.Ansible,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d, constant.RunsTable)
	s := store.New(d, constant.RunsTable)
	s.AbandonRunning()
	p := reaper.New(r)
	result := &Service{
		database: d,
		store:    s,
		reaper:   p,
		runner:   runner.New(o, s, l, r, p),
	}

	if o.Inventory != nil {
		result.inventory = o.Inventory()
		result.runner.WithInventory(result.inventory)
	}

	return result
}
