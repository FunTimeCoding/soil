package toolchain

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func mountedIndex(
	t *testing.T,
	directory string,
) *client.Client {
	t.Helper()
	d, _ := daemon_tester.NewInIndex(t, inventory.New(), directory)
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result := client.New(
		locator.New("localhost").Port(s.Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
