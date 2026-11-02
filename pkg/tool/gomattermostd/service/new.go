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
		logger:     l,
		reporter:   r,
		debounce:   o.DebounceWindow,
		purge:      o.PurgeWindow,
	}

	if result.connector != nil {
		result.rosterGuard = peer.New(
			constant.RosterPeer,
			result.connector.Address(),
			result.connector,
			l,
			r,
		)
		result.notifier = notifier.New(
			result.connector,
			constant.Identity.Name(),
			result.rosterGuard,
		)
		result.watcher = watcher.New(m, g, result.store, l, r)
	}

	if o.Monitor.Enabled {
		result.monitor = monitor.New(m, g, o.Monitor, l)
	}

	return result
}
