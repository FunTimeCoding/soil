package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
	"time"
)

func webServer(t *testing.T) (*model_context_server.Server, *store.Store) {
	t.Helper()
	l := logger.New(t.Context())
	o := store_tester.NewOption(t, time.Now, notifier.New())
	r := memory.New()
	d := daemon.New(o, l, r)
	t.Cleanup(d.Close)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(result.Close)
	m := lite.New(l, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })

	return result, store.New(m, time.Now)
}
