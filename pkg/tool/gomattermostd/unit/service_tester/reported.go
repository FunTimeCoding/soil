package service_tester

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"testing"
)

func Reported(
	t *testing.T,
	c chat.MattermostClient,
	k goclauded.Client,
	configure func(*option.Mattermost),
	r soilFace.Reporter,
) *service.Service {
	t.Helper()
	o := NewOption(t, c, k, configure)
	l := logger.New(t.Context())
	m := lite.New(l, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })
	migrate.AutoMigrate(m)
	s := store.New(m)
	g := peer.New(constant.MattermostPeer, c.Host(), c, l, r)
	t.Cleanup(g.Close)
	result := service.New(s, c, g, l, r, o.DebounceWindow, o.PurgeWindow)

	if k != nil {
		h := peer.New(constant.RosterPeer, k.Address(), k, l, r)
		t.Cleanup(h.Close)
		w := watcher.New(c, g, s, l, r).WithBackoff(
			o.ReconnectMinimum,
			o.ReconnectMaximum,
		)
		t.Cleanup(w.Close)
		result.WithSubscription(notifier.New(k, constant.Identity.Name(), h), w)
	}

	if o.Monitor.Enabled {
		result.WithMonitor(monitor.New(c, g, o.Monitor, l))
	}

	return result
}
