package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.RequiredInteger(constant.PortEnvironment),
		environment.Exists(constant.InsecureEnvironment),
	)
}
