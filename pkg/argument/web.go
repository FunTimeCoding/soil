package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func (i *Instance) Web() {
	i.Integer(
		constant.Port,
		environment.FallbackInteger(
			webConstant.PortEnvironment,
			webConstant.ListenPort,
		),
		webConstant.PortUsage,
	)
	i.String(
		constant.BindAddress,
		environment.Fallback(webConstant.BindEnvironment, webConstant.Loopback),
		webConstant.BindUsage,
	)
}
