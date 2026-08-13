package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	web "github.com/funtimecoding/soil/pkg/web/constant"
)

func (i *Instance) Web() {
	i.Integer(
		constant.Port,
		environment.FallbackInteger(web.PortEnvironment, web.ListenPort),
		web.PortUsage,
	)
	i.String(
		constant.BindAddress,
		environment.Fallback(web.BindEnvironment, web.Loopback),
		web.BindUsage,
	)
}
