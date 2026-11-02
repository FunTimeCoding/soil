package client

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	goclaudeConstant "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func NewEnvironment() *Client {
	return New(
		locator.Environment(
			goclaudeConstant.HostEnvironment,
			goclaudeConstant.PortEnvironment,
			goclaudeConstant.InsecureEnvironment,
		).String(),
		environment.Exists(goclaudeConstant.UntrustedEnvironment),
		environment.Required(goclaudeConstant.TokenEnvironment),
	)
}
