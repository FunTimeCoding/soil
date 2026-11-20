package monitor_tester

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	gomonitordConstant "github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"strconv"
)

func (o *Tester) Connect(port int) {
	o.t.Helper()
	o.t.Setenv(gomonitordConstant.HostEnvironment, webConstant.Localhost)
	o.t.Setenv(gomonitordConstant.PortEnvironment, strconv.Itoa(port))
	o.t.Setenv(gomonitordConstant.InsecureEnvironment, "true")
	o.t.Setenv(
		gomonitordConstant.TokenEnvironment,
		generativeConstant.ModelContextTestToken,
	)
}
