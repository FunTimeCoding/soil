package connector

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewEnvironment() *Client {
	return New(
		locator.Environment(
			constant.HostEnvironment,
			constant.PortEnvironment,
			constant.InsecureEnvironment,
		).String(),
		environment.Exists(constant.UntrustedEnvironment),
		environment.Required(constant.TokenEnvironment),
	)
}
