package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
)

func NewEnvironment(instance string) *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.RequiredInteger(constant.PortEnvironment),
		environment.Exists(constant.InsecureEnvironment),
		instance,
	)
}
