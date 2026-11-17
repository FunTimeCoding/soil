package unit

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func TestGuard(t *testing.T) {
	d := service_tester.New(t).Daemon()
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(g)
		},
		model_context_server.WithSignOn(),
	)
	defer s.Close()
	s.VerifyBase(t)
	s.VerifyFavicon(t)
	s.VerifySignIn(t)
}
