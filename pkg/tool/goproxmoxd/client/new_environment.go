package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewEnvironment(instance string) *Client {
	return New(
		locator.Environment(
			constant.HostEnvironment,
			constant.PortEnvironment,
			constant.InsecureEnvironment,
		),
		instance,
		environment.Required(constant.TokenEnvironment),
	)
}
