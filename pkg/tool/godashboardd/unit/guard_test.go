package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/authorization/authorization_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func TestGuard(t *testing.T) {
	d := daemon.New(
		service_tester.NewOption(
			t,
			service_tester.FixtureBoard(),
			newPrometheus(t),
			nil,
			nil,
		),
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(d.Close)
	s := model_context_server.NewSignedOn(
		t,
		authorization_tester.Client,
		func(g *guard.Mux) {
			d.Mount(g)
		},
	)
	defer s.Close()
	s.VerifyBase(t)
	s.VerifyFavicon(t)
	s.VerifySignIn(t)
}
