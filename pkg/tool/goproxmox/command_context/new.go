package command_context

import (
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/client"
)

func New(t *terminal.Terminal) *Context {
	return &Context{client: client.NewEnvironment(), terminal: t}
}
