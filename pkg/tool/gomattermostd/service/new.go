package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/sweeper"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
)

func New(
	o *option.Mattermost,
	l *logger.Logger,
	r face.Reporter,
) *Service {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	m := o.Mattermost()
	g := peer.New(constant.MattermostPeer, m.Host(), m, l, r)
	result := &Service{
		database:   d,
		store:      store.New(d),
		mattermost: m,
		guard:      g,
		connector:  o.Connector(),
	}

	if result.connector != nil {
		n := notifier.New(result.connector, constant.Identity.Name(), r)
		result.watcher = watcher.New(
			m,
			g,
			result.store,
			n,
			l,
			r,
			o.DebounceWindow,
		)
		result.sweeper = sweeper.New(
			result.store,
			n,
			result.watcher,
			l,
			r,
			o.PurgeWindow,
			o.PurgeInterval,
		)
	}

	if o.Monitor.Enabled {
		result.monitor = monitor.New(m, g, o.Monitor, l, r)
	}

	return result
}
