package guard

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	c := mock_client.New()
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			golinkaced.Mount(
				c,
				service.New(c),
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
				g,
			)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyInterface(t)
	v.VerifyGuarded(t, "/api/links")
	v.VerifyGuarded(t, "/api/search")
	v.VerifyModelContext(t)
}
