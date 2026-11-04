package subscribed

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"testing"
	"time"
)

func New(t *testing.T) *Subscribed {
	t.Helper()
	r := memory.New()
	l := logger.New(t.Context())
	m := service_tester.NewMattermost(t, "selfuser")
	s := worker_tester.NewStore(t)
	g := worker_tester.NewGuard(t, m, r)
	k := mock_client.New()
	h := peer.New(constant.RosterPeer, k.Address(), k, l, r)
	t.Cleanup(h.Close)
	w := watcher.New(m, g, s, l, r)
	t.Cleanup(w.Close)
	result := &Subscribed{
		Service: service.New(s, m, g, l, r, 0, time.Hour).WithSubscription(
			notifier.New(k, constant.Identity.Name(), h),
			w,
		),
		Watcher:    w,
		Mattermost: m,
		Roster:     k,
	}
	t.Cleanup(result.Close)

	return result
}
