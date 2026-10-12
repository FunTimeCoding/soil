package telemetry

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func NewEnvironment() *Client {
	return New(
		environment.Fallback(constant.HostEnvironment, webConstant.Localhost),
		environment.FallbackInteger(constant.PortEnvironment, 0),
		environment.Exists(constant.InsecureEnvironment),
	)
}
