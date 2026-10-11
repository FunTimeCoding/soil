package unit

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	monitorConstant "github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"strconv"
	"testing"
)

func connectTo(
	t *testing.T,
	port int,
) {
	t.Helper()
	t.Setenv(monitorConstant.HostEnvironment, webConstant.Localhost)
	t.Setenv(monitorConstant.PortEnvironment, strconv.Itoa(port))
	t.Setenv(monitorConstant.InsecureEnvironment, "true")
	t.Setenv(monitorConstant.TokenEnvironment, constant.ModelContextTestToken)
}
