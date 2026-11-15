package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	j := service_tester.NewJira(t)
	c := service_tester.NewConfluence(t)

	switch peer {
	case constant.JiraPeer:
		j.Failure = outage_tester.RefusedDial()
	case constant.ConfluencePeer:
		c.Failure = outage_tester.RefusedDial()
	}

	d := service_tester.NewDaemon(t, service_tester.NewOption(j, c, nil), r)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(t, d.Workers(), peer).Unavailable
}
