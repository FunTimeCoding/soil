package unit

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	gomonitordConstant "github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"strconv"
	"testing"
)

func connectTo(
	t *testing.T,
	port int,
) {
	t.Helper()
	t.Setenv(gomonitordConstant.HostEnvironment, webConstant.Localhost)
	t.Setenv(gomonitordConstant.PortEnvironment, strconv.Itoa(port))
	t.Setenv(gomonitordConstant.InsecureEnvironment, "true")
	t.Setenv(
		gomonitordConstant.TokenEnvironment,
		generativeConstant.ModelContextTestToken,
	)
}
