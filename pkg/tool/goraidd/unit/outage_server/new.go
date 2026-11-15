package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	goraidparsedConstant "github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	p := service_tester.NewParser(t)
	p.Fail(outage_tester.RefusedDial())
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(t, t.TempDir(), t.TempDir(), p),
		r,
	)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
		model_context_server.WithSignOn(),
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(
		t,
		d.Workers(),
		goraidparsedConstant.Identity.Name(),
	).Unavailable
}
