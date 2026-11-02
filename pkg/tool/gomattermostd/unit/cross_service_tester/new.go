package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"net/http"
	"path/filepath"
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
	t.Cleanup(claude.Close)
	upstream := mattermost_client_tester.New(t, configure)
	r := memory.New()
	path := filepath.Join(t.TempDir(), "mattermost.db")
	s := worker_tester.NewStoreOnFile(t, path)
	v := service_tester.Reported(
		t,
		upstream.Client,
		claude.Client(t),
		func(o *option.Mattermost) {
			o.LitePath = path
			o.DebounceWindow = window
		},
		r,
	)

	return &Tester{
		t:         t,
		Upstream:  upstream,
		Store:     s,
		Service:   v,
		Goclauded: claude,
		Reporter:  r,
	}
}
