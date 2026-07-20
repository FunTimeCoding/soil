package callback

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/authorization/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func NewEnvironment(verbose bool) *Server {
	return New(
		environment.FallbackInteger(
			constant.PortEnvironment,
			webConstant.ListenPort,
		),
		verbose,
	)
}
