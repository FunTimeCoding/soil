package cross_service_tester

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/chat/unit/mattermost_client_tester"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
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
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "mattermost.db")
	o.Mattermost = func() chat.MattermostClient { return upstream.Client }
	k := claude.Client(t)
	o.Connector = func() face.Client { return k }
	o.DebounceWindow = window
	o.PurgeWindow = 365 * 24 * time.Hour
	o.FlushInterval = constant.FixtureFlushInterval
	o.ReconnectMinimum = constant.FixtureReconnectMinimum
	o.ReconnectMaximum = constant.FixtureReconnectMaximum
	d := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(d.Close)

	return &Tester{
		t:         t,
		Upstream:  upstream,
		Store:     worker_tester.NewStoreOnFile(t, o.LitePath),
		Daemon:    d,
		Goclauded: claude,
		Reporter:  r,
	}
}
