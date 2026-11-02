package telemetry

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	gotelemetrydConstant "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func NewEnvironment() *Client {
	return New(
		environment.Fallback(
			gotelemetrydConstant.HostEnvironment,
			webConstant.Localhost,
		),
		environment.FallbackInteger(gotelemetrydConstant.PortEnvironment, 0),
		environment.Exists(gotelemetrydConstant.InsecureEnvironment),
	)
}
