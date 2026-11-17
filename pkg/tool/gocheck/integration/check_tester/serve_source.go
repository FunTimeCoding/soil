package check_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	gosourcedConstant "github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/daemon_tester"
	"strconv"
	"testing"
)

func ServeSource(
	t *testing.T,
	i *inventory.Inventory,
) {
	t.Helper()
	d, _ := daemon_tester.NewInIndex(t, i, t.TempDir())
	s := model_context_server.Mount(t, d, memory.New())
	t.Setenv(gosourcedConstant.HostEnvironment, "localhost")
	t.Setenv(gosourcedConstant.PortEnvironment, strconv.Itoa(s.Port))
	t.Setenv(gosourcedConstant.InsecureEnvironment, "true")
	t.Setenv(
		gosourcedConstant.TokenEnvironment,
		generativeConstant.ModelContextTestToken,
	)
}
