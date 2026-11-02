package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	s := service_tester.New(t)
	d := daemon.New(
		s.Option,
		logger.New(context.Background()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/placements")
	v.VerifyGuarded(t, "/api/places")
	v.VerifyGuarded(t, "/api/sightings")
	v.VerifyOpen(t, "/favicon.ico")
	v.VerifySignIn(t)
}
