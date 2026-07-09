package argument

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func (i *Instance) Web() {
	i.Integer(
		Port,
		environment.FallbackInteger(constant.PortEnvironment, constant.ListenPort),
		constant.PortUsage,
	)
	i.String(
		BindAddress,
		environment.Fallback(constant.BindEnvironment, constant.Loopback),
		constant.BindUsage,
	)
}
