package telemetry

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/telemetry/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
)

func NewEnvironment() *Client {
	return New(
		environment.Fallback(constant.HostEnvironment, web.Localhost),
		environment.FallbackInteger(constant.PortEnvironment, web.ListenPort),
		environment.Exists(constant.InsecureEnvironment),
	)
}
