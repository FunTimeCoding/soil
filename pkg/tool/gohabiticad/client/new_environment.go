package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewEnvironment() *Client {
	return New(
		locator.Environment(
			constant.HostEnvironment,
			constant.PortEnvironment,
			constant.InsecureEnvironment,
		),
		environment.Required(constant.TokenEnvironment),
	)
}
