package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/integration/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"net/http"
	"testing"
	"time"
)

func New(
	t *testing.T,
	configure func(*http.ServeMux),
	window time.Duration,
) *Tester {
	t.Helper()
	claude := base.New(t)
	upstream := mattermost_client_tester.New(t, configure)
	s := store.New(lite.NewMemory())
	r := memory.New()
	w := watcher.New(
		upstream.Client,
		s,
		notifier.New(claude.Connector(t), "mattermost", r),
		logger.New(t.Context()),
		r,
		window,
	)
	t.Cleanup(
		func() {
			w.Stop()
			s.Close()
			claude.Stop()
		},
	)

	return &Tester{
		Upstream:  upstream,
		Store:     s,
		Watcher:   w,
		Goclauded: claude,
		Reporter:  r,
	}
}
