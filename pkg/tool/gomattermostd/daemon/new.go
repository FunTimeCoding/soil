package daemon

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
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"github.com/funtimecoding/soil/pkg/worker"
)

func New(
	o *option.Mattermost,
	l *logger.Logger,
	r face.Reporter,
) *Daemon {
	d := relational.Open(l, o.PostgresLocator, o.LitePath)
	migrate.AutoMigrate(d)
	s := store.New(d)
	m := o.Mattermost()
	g := peer.New(constant.MattermostPeer, m.Host(), m, l, r)
	v := service.New(s, m, g, l, r, o.DebounceWindow, o.PurgeWindow)
	result := &Daemon{
		database:   d,
		mattermost: m,
		guard:      g,
		service:    v,
	}

	if c := o.Connector(); c != nil {
		result.connector = c
		result.rosterGuard = peer.New(constant.RosterPeer, c.Address(), c, l, r)
		result.watcher = watcher.New(m, g, s, l, r).WithBackoff(
			o.ReconnectMinimum,
			o.ReconnectMaximum,
		)
		v.WithSubscription(
			notifier.New(c, constant.Identity.Name(), result.rosterGuard),
			result.watcher,
		)
		result.clocks = append(
			result.clocks,
			worker.NewEvery(v.FlushDue, l, r, o.FlushInterval),
			worker.NewEvery(v.Purge, l, r, o.PurgeInterval),
		)
	}

	if o.Monitor.Enabled {
		v.WithMonitor(monitor.New(m, g, o.Monitor, l))
		result.clocks = append(
			result.clocks,
			worker.NewEvery(v.PollMonitor, l, r, o.Monitor.Interval),
		)
	}

	return result
}
