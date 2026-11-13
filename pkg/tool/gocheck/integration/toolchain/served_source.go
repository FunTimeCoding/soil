package toolchain

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	gosourcedConstant "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"strconv"
	"testing"
)

func servedSource(
	t *testing.T,
	i *inventory.Inventory,
) {
	t.Helper()
	d, _ := daemon_tester.NewInIndex(t, i, t.TempDir())
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	t.Setenv(gosourcedConstant.HostEnvironment, "localhost")
	t.Setenv(gosourcedConstant.PortEnvironment, strconv.Itoa(s.Port))
	t.Setenv(gosourcedConstant.InsecureEnvironment, "true")
	t.Setenv(
		gosourcedConstant.TokenEnvironment,
		generativeConstant.ModelContextTestToken,
	)
}
