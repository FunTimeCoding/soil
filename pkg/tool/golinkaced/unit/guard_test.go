package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	s := service_tester.Service(t, service_tester.NewSource(t), memory.New())
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			golinkaced.Mount(s, memory.New(), mock_recorder.New(), g)
		},
	)
	defer v.Close()
	v.VerifyBase(t)
	v.VerifyGuarded(t, "/api/links")
	v.VerifyGuarded(t, "/api/search")
	v.VerifyModelContext(t)
}
