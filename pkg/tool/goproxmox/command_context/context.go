package command_context

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/client"
)

type Context struct {
	client   *client.Client
	terminal *terminal.Terminal
}
